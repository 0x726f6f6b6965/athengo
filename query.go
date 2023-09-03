package athengo

import (
	"context"
	"errors"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/athena"
)

// startQuery starts an Athena query and returns its ID.
func (c *athenaClient) startQuery(query string) (string, error) {
	resp, err := c.athena.StartQueryExecution(&athena.StartQueryExecutionInput{
		QueryString: aws.String(query),
		QueryExecutionContext: &athena.QueryExecutionContext{
			Database: aws.String(c.database),
		},
		ResultConfiguration: &athena.ResultConfiguration{
			OutputLocation: aws.String(c.outputLocation),
		},
	})
	if err != nil {
		return "", err
	}

	return *resp.QueryExecutionId, nil
}

// waitOnQuery blocks until a query finishes, returning an error if it failed.
func (c *athenaClient) waitOnQuery(ctx context.Context, queryID string) error {
	for {
		statusResp, err := c.athena.GetQueryExecutionWithContext(ctx, &athena.GetQueryExecutionInput{
			QueryExecutionId: aws.String(queryID),
		})
		if err != nil {
			return err
		}

		switch *statusResp.QueryExecution.Status.State {
		case athena.QueryExecutionStateCancelled:
			return context.Canceled
		case athena.QueryExecutionStateFailed:
			reason := *statusResp.QueryExecution.Status.StateChangeReason
			return errors.New(reason)
		case athena.QueryExecutionStateSucceeded:
			return nil
		case athena.QueryExecutionStateQueued:
		case athena.QueryExecutionStateRunning:
		}

		select {
		case <-ctx.Done():
			c.athena.StopQueryExecution(&athena.StopQueryExecutionInput{
				QueryExecutionId: aws.String(queryID),
			})

			return ctx.Err()
		case <-time.After(c.pollFrequency):
			continue
		}
	}
}

func (c *athenaClient) ExecQuery(ctx context.Context, query string) (Rows, error) {
	queryId, err := c.startQuery(query)
	if err != nil {
		return nil, err
	}

	if waitErr := c.waitOnQuery(ctx, queryId); waitErr != nil {
		return nil, err
	}

	r, err := NewRows(c.GetAthenaApi(), queryId)
	if err != nil {
		return nil, err
	}
	return r, nil
}
