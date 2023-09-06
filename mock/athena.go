package mock

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/athena"
	"github.com/aws/aws-sdk-go/service/athena/athenaiface"
)

type mockAthena struct {
}

type Key int

const (
	STATUS Key = 0
)

// BatchGetNamedQuery implements athenaiface.AthenaAPI.
func (*mockAthena) BatchGetNamedQuery(*athena.BatchGetNamedQueryInput) (*athena.BatchGetNamedQueryOutput, error) {
	panic("unimplemented")
}

// BatchGetNamedQueryRequest implements athenaiface.AthenaAPI.
func (*mockAthena) BatchGetNamedQueryRequest(*athena.BatchGetNamedQueryInput) (*request.Request, *athena.BatchGetNamedQueryOutput) {
	panic("unimplemented")
}

// BatchGetNamedQueryWithContext implements athenaiface.AthenaAPI.
func (*mockAthena) BatchGetNamedQueryWithContext(context.Context, *athena.BatchGetNamedQueryInput, ...request.Option) (*athena.BatchGetNamedQueryOutput, error) {
	panic("unimplemented")
}

// BatchGetQueryExecution implements athenaiface.AthenaAPI.
func (*mockAthena) BatchGetQueryExecution(*athena.BatchGetQueryExecutionInput) (*athena.BatchGetQueryExecutionOutput, error) {
	panic("unimplemented")
}

// BatchGetQueryExecutionRequest implements athenaiface.AthenaAPI.
func (*mockAthena) BatchGetQueryExecutionRequest(*athena.BatchGetQueryExecutionInput) (*request.Request, *athena.BatchGetQueryExecutionOutput) {
	panic("unimplemented")
}

// BatchGetQueryExecutionWithContext implements athenaiface.AthenaAPI.
func (*mockAthena) BatchGetQueryExecutionWithContext(context.Context, *athena.BatchGetQueryExecutionInput, ...request.Option) (*athena.BatchGetQueryExecutionOutput, error) {
	panic("unimplemented")
}

// CreateNamedQuery implements athenaiface.AthenaAPI.
func (*mockAthena) CreateNamedQuery(*athena.CreateNamedQueryInput) (*athena.CreateNamedQueryOutput, error) {
	panic("unimplemented")
}

// CreateNamedQueryRequest implements athenaiface.AthenaAPI.
func (*mockAthena) CreateNamedQueryRequest(*athena.CreateNamedQueryInput) (*request.Request, *athena.CreateNamedQueryOutput) {
	panic("unimplemented")
}

// CreateNamedQueryWithContext implements athenaiface.AthenaAPI.
func (*mockAthena) CreateNamedQueryWithContext(context.Context, *athena.CreateNamedQueryInput, ...request.Option) (*athena.CreateNamedQueryOutput, error) {
	panic("unimplemented")
}

// CreateWorkGroup implements athenaiface.AthenaAPI.
func (*mockAthena) CreateWorkGroup(*athena.CreateWorkGroupInput) (*athena.CreateWorkGroupOutput, error) {
	panic("unimplemented")
}

// CreateWorkGroupRequest implements athenaiface.AthenaAPI.
func (*mockAthena) CreateWorkGroupRequest(*athena.CreateWorkGroupInput) (*request.Request, *athena.CreateWorkGroupOutput) {
	panic("unimplemented")
}

// CreateWorkGroupWithContext implements athenaiface.AthenaAPI.
func (*mockAthena) CreateWorkGroupWithContext(context.Context, *athena.CreateWorkGroupInput, ...request.Option) (*athena.CreateWorkGroupOutput, error) {
	panic("unimplemented")
}

// DeleteNamedQuery implements athenaiface.AthenaAPI.
func (*mockAthena) DeleteNamedQuery(*athena.DeleteNamedQueryInput) (*athena.DeleteNamedQueryOutput, error) {
	panic("unimplemented")
}

// DeleteNamedQueryRequest implements athenaiface.AthenaAPI.
func (*mockAthena) DeleteNamedQueryRequest(*athena.DeleteNamedQueryInput) (*request.Request, *athena.DeleteNamedQueryOutput) {
	panic("unimplemented")
}

// DeleteNamedQueryWithContext implements athenaiface.AthenaAPI.
func (*mockAthena) DeleteNamedQueryWithContext(context.Context, *athena.DeleteNamedQueryInput, ...request.Option) (*athena.DeleteNamedQueryOutput, error) {
	panic("unimplemented")
}

// DeleteWorkGroup implements athenaiface.AthenaAPI.
func (*mockAthena) DeleteWorkGroup(*athena.DeleteWorkGroupInput) (*athena.DeleteWorkGroupOutput, error) {
	panic("unimplemented")
}

// DeleteWorkGroupRequest implements athenaiface.AthenaAPI.
func (*mockAthena) DeleteWorkGroupRequest(*athena.DeleteWorkGroupInput) (*request.Request, *athena.DeleteWorkGroupOutput) {
	panic("unimplemented")
}

// DeleteWorkGroupWithContext implements athenaiface.AthenaAPI.
func (*mockAthena) DeleteWorkGroupWithContext(context.Context, *athena.DeleteWorkGroupInput, ...request.Option) (*athena.DeleteWorkGroupOutput, error) {
	panic("unimplemented")
}

// GetNamedQuery implements athenaiface.AthenaAPI.
func (*mockAthena) GetNamedQuery(*athena.GetNamedQueryInput) (*athena.GetNamedQueryOutput, error) {
	panic("unimplemented")
}

// GetNamedQueryRequest implements athenaiface.AthenaAPI.
func (*mockAthena) GetNamedQueryRequest(*athena.GetNamedQueryInput) (*request.Request, *athena.GetNamedQueryOutput) {
	panic("unimplemented")
}

// GetNamedQueryWithContext implements athenaiface.AthenaAPI.
func (*mockAthena) GetNamedQueryWithContext(context.Context, *athena.GetNamedQueryInput, ...request.Option) (*athena.GetNamedQueryOutput, error) {
	panic("unimplemented")
}

// GetQueryExecution implements athenaiface.AthenaAPI.
func (*mockAthena) GetQueryExecution(*athena.GetQueryExecutionInput) (*athena.GetQueryExecutionOutput, error) {
	panic("unimplemented")
}

// GetQueryExecutionRequest implements athenaiface.AthenaAPI.
func (*mockAthena) GetQueryExecutionRequest(*athena.GetQueryExecutionInput) (*request.Request, *athena.GetQueryExecutionOutput) {
	panic("unimplemented")
}

// GetQueryExecutionWithContext implements athenaiface.AthenaAPI.
func (*mockAthena) GetQueryExecutionWithContext(ctx context.Context, query *athena.GetQueryExecutionInput, params ...request.Option) (*athena.GetQueryExecutionOutput, error) {
	result := &athena.GetQueryExecutionOutput{

		QueryExecution: &athena.QueryExecution{
			QueryExecutionId: aws.String("test"),
			Status:           &athena.QueryExecutionStatus{State: aws.String(athena.QueryExecutionStateSucceeded)},
		},
	}

	if val, ok := ctx.Value(STATUS).(string); ok {
		switch val {
		case athena.QueryExecutionStateFailed:
			result.QueryExecution.Status.StateChangeReason = aws.String("Test Failed")
		}
		result.QueryExecution.Status.State = aws.String(val)
	}
	return result, nil
}

// GetQueryResults implements athenaiface.AthenaAPI.
func (*mockAthena) GetQueryResults(*athena.GetQueryResultsInput) (*athena.GetQueryResultsOutput, error) {
	return &athena.GetQueryResultsOutput{
		ResultSet: &athena.ResultSet{
			ResultSetMetadata: &athena.ResultSetMetadata{
				ColumnInfo: []*athena.ColumnInfo{
					{
						CaseSensitive: aws.Bool(false),
						CatalogName:   aws.String("hive"),
						Label:         aws.String("id"),
						Name:          aws.String("id"),
						Nullable:      aws.String("UNKNOWN"),
						Precision:     aws.Int64(19),
						Scale:         aws.Int64(0),
						Type:          aws.String("bigint"),
					},
					{
						CaseSensitive: aws.Bool(true),
						CatalogName:   aws.String("hive"),
						Label:         aws.String("name"),
						Name:          aws.String("name"),
						Nullable:      aws.String("UNKNOWN"),
						Precision:     aws.Int64(2147483647),
						Scale:         aws.Int64(0),
						Type:          aws.String("varchar"),
					},
					{
						CaseSensitive: aws.Bool(false),
						CatalogName:   aws.String("hive"),
						Label:         aws.String("create_time"),
						Name:          aws.String("create_time"),
						Nullable:      aws.String("UNKNOWN"),
						Precision:     aws.Int64(3),
						Scale:         aws.Int64(0),
						Type:          aws.String("timestamp"),
					},
					{
						CaseSensitive: aws.Bool(false),
						CatalogName:   aws.String("hive"),
						Label:         aws.String("update_time"),
						Name:          aws.String("update_time"),
						Nullable:      aws.String("UNKNOWN"),
						Precision:     aws.Int64(3),
						Scale:         aws.Int64(0),
						Type:          aws.String("timestamp"),
					},
				},
			},
			Rows: []*athena.Row{
				{
					Data: []*athena.Datum{
						{
							VarCharValue: aws.String("id"),
						},
						{
							VarCharValue: aws.String("name"),
						},
						{
							VarCharValue: aws.String("create_time"),
						},
						{
							VarCharValue: aws.String("update_time"),
						},
					},
				},
				{
					Data: []*athena.Datum{
						{
							VarCharValue: aws.String("1"),
						},
						{
							VarCharValue: aws.String("A"),
						},
						{
							VarCharValue: aws.String("2023-08-09 00:04:01.000"),
						},
						{
							VarCharValue: aws.String("2023-08-10 00:04:01.000"),
						},
					},
				},
				{
					Data: []*athena.Datum{
						{
							VarCharValue: aws.String("2"),
						},
						{
							VarCharValue: aws.String("B"),
						},
						{
							VarCharValue: aws.String("2023-08-12 00:04:01.000"),
						},
						{
							VarCharValue: aws.String("2023-08-12 00:04:01.000"),
						},
					},
				},
				{
					Data: []*athena.Datum{
						{
							VarCharValue: aws.String("3"),
						},
						{
							VarCharValue: aws.String("C"),
						},
						{
							VarCharValue: aws.String("2023-08-13 00:04:01.000"),
						},
						{
							VarCharValue: aws.String("2023-08-13 00:04:01.000"),
						},
					},
				},
			},
		},
	}, nil
}

// GetQueryResultsPages implements athenaiface.AthenaAPI.
func (*mockAthena) GetQueryResultsPages(*athena.GetQueryResultsInput, func(*athena.GetQueryResultsOutput, bool) bool) error {
	panic("unimplemented")
}

// GetQueryResultsPagesWithContext implements athenaiface.AthenaAPI.
func (*mockAthena) GetQueryResultsPagesWithContext(context.Context, *athena.GetQueryResultsInput, func(*athena.GetQueryResultsOutput, bool) bool, ...request.Option) error {
	panic("unimplemented")
}

// GetQueryResultsRequest implements athenaiface.AthenaAPI.
func (*mockAthena) GetQueryResultsRequest(*athena.GetQueryResultsInput) (*request.Request, *athena.GetQueryResultsOutput) {
	panic("unimplemented")
}

// GetQueryResultsWithContext implements athenaiface.AthenaAPI.
func (*mockAthena) GetQueryResultsWithContext(context.Context, *athena.GetQueryResultsInput, ...request.Option) (*athena.GetQueryResultsOutput, error) {
	panic("unimplemented")
}

// GetWorkGroup implements athenaiface.AthenaAPI.
func (*mockAthena) GetWorkGroup(*athena.GetWorkGroupInput) (*athena.GetWorkGroupOutput, error) {
	panic("unimplemented")
}

// GetWorkGroupRequest implements athenaiface.AthenaAPI.
func (*mockAthena) GetWorkGroupRequest(*athena.GetWorkGroupInput) (*request.Request, *athena.GetWorkGroupOutput) {
	panic("unimplemented")
}

// GetWorkGroupWithContext implements athenaiface.AthenaAPI.
func (*mockAthena) GetWorkGroupWithContext(context.Context, *athena.GetWorkGroupInput, ...request.Option) (*athena.GetWorkGroupOutput, error) {
	panic("unimplemented")
}

// ListNamedQueries implements athenaiface.AthenaAPI.
func (*mockAthena) ListNamedQueries(*athena.ListNamedQueriesInput) (*athena.ListNamedQueriesOutput, error) {
	panic("unimplemented")
}

// ListNamedQueriesPages implements athenaiface.AthenaAPI.
func (*mockAthena) ListNamedQueriesPages(*athena.ListNamedQueriesInput, func(*athena.ListNamedQueriesOutput, bool) bool) error {
	panic("unimplemented")
}

// ListNamedQueriesPagesWithContext implements athenaiface.AthenaAPI.
func (*mockAthena) ListNamedQueriesPagesWithContext(context.Context, *athena.ListNamedQueriesInput, func(*athena.ListNamedQueriesOutput, bool) bool, ...request.Option) error {
	panic("unimplemented")
}

// ListNamedQueriesRequest implements athenaiface.AthenaAPI.
func (*mockAthena) ListNamedQueriesRequest(*athena.ListNamedQueriesInput) (*request.Request, *athena.ListNamedQueriesOutput) {
	panic("unimplemented")
}

// ListNamedQueriesWithContext implements athenaiface.AthenaAPI.
func (*mockAthena) ListNamedQueriesWithContext(context.Context, *athena.ListNamedQueriesInput, ...request.Option) (*athena.ListNamedQueriesOutput, error) {
	panic("unimplemented")
}

// ListQueryExecutions implements athenaiface.AthenaAPI.
func (*mockAthena) ListQueryExecutions(*athena.ListQueryExecutionsInput) (*athena.ListQueryExecutionsOutput, error) {
	panic("unimplemented")
}

// ListQueryExecutionsPages implements athenaiface.AthenaAPI.
func (*mockAthena) ListQueryExecutionsPages(*athena.ListQueryExecutionsInput, func(*athena.ListQueryExecutionsOutput, bool) bool) error {
	panic("unimplemented")
}

// ListQueryExecutionsPagesWithContext implements athenaiface.AthenaAPI.
func (*mockAthena) ListQueryExecutionsPagesWithContext(context.Context, *athena.ListQueryExecutionsInput, func(*athena.ListQueryExecutionsOutput, bool) bool, ...request.Option) error {
	panic("unimplemented")
}

// ListQueryExecutionsRequest implements athenaiface.AthenaAPI.
func (*mockAthena) ListQueryExecutionsRequest(*athena.ListQueryExecutionsInput) (*request.Request, *athena.ListQueryExecutionsOutput) {
	panic("unimplemented")
}

// ListQueryExecutionsWithContext implements athenaiface.AthenaAPI.
func (*mockAthena) ListQueryExecutionsWithContext(context.Context, *athena.ListQueryExecutionsInput, ...request.Option) (*athena.ListQueryExecutionsOutput, error) {
	panic("unimplemented")
}

// ListTagsForResource implements athenaiface.AthenaAPI.
func (*mockAthena) ListTagsForResource(*athena.ListTagsForResourceInput) (*athena.ListTagsForResourceOutput, error) {
	panic("unimplemented")
}

// ListTagsForResourceRequest implements athenaiface.AthenaAPI.
func (*mockAthena) ListTagsForResourceRequest(*athena.ListTagsForResourceInput) (*request.Request, *athena.ListTagsForResourceOutput) {
	panic("unimplemented")
}

// ListTagsForResourceWithContext implements athenaiface.AthenaAPI.
func (*mockAthena) ListTagsForResourceWithContext(context.Context, *athena.ListTagsForResourceInput, ...request.Option) (*athena.ListTagsForResourceOutput, error) {
	panic("unimplemented")
}

// ListWorkGroups implements athenaiface.AthenaAPI.
func (*mockAthena) ListWorkGroups(*athena.ListWorkGroupsInput) (*athena.ListWorkGroupsOutput, error) {
	panic("unimplemented")
}

// ListWorkGroupsPages implements athenaiface.AthenaAPI.
func (*mockAthena) ListWorkGroupsPages(*athena.ListWorkGroupsInput, func(*athena.ListWorkGroupsOutput, bool) bool) error {
	panic("unimplemented")
}

// ListWorkGroupsPagesWithContext implements athenaiface.AthenaAPI.
func (*mockAthena) ListWorkGroupsPagesWithContext(context.Context, *athena.ListWorkGroupsInput, func(*athena.ListWorkGroupsOutput, bool) bool, ...request.Option) error {
	panic("unimplemented")
}

// ListWorkGroupsRequest implements athenaiface.AthenaAPI.
func (*mockAthena) ListWorkGroupsRequest(*athena.ListWorkGroupsInput) (*request.Request, *athena.ListWorkGroupsOutput) {
	panic("unimplemented")
}

// ListWorkGroupsWithContext implements athenaiface.AthenaAPI.
func (*mockAthena) ListWorkGroupsWithContext(context.Context, *athena.ListWorkGroupsInput, ...request.Option) (*athena.ListWorkGroupsOutput, error) {
	panic("unimplemented")
}

// StartQueryExecution implements athenaiface.AthenaAPI.
func (*mockAthena) StartQueryExecution(*athena.StartQueryExecutionInput) (*athena.StartQueryExecutionOutput, error) {
	return &athena.StartQueryExecutionOutput{
		QueryExecutionId: aws.String("test"),
	}, nil
}

// StartQueryExecutionRequest implements athenaiface.AthenaAPI.
func (*mockAthena) StartQueryExecutionRequest(*athena.StartQueryExecutionInput) (*request.Request, *athena.StartQueryExecutionOutput) {
	panic("unimplemented")
}

// StartQueryExecutionWithContext implements athenaiface.AthenaAPI.
func (*mockAthena) StartQueryExecutionWithContext(context.Context, *athena.StartQueryExecutionInput, ...request.Option) (*athena.StartQueryExecutionOutput, error) {
	panic("unimplemented")
}

// StopQueryExecution implements athenaiface.AthenaAPI.
func (*mockAthena) StopQueryExecution(*athena.StopQueryExecutionInput) (*athena.StopQueryExecutionOutput, error) {
	panic("unimplemented")
}

// StopQueryExecutionRequest implements athenaiface.AthenaAPI.
func (*mockAthena) StopQueryExecutionRequest(*athena.StopQueryExecutionInput) (*request.Request, *athena.StopQueryExecutionOutput) {
	panic("unimplemented")
}

// StopQueryExecutionWithContext implements athenaiface.AthenaAPI.
func (*mockAthena) StopQueryExecutionWithContext(context.Context, *athena.StopQueryExecutionInput, ...request.Option) (*athena.StopQueryExecutionOutput, error) {
	panic("unimplemented")
}

// TagResource implements athenaiface.AthenaAPI.
func (*mockAthena) TagResource(*athena.TagResourceInput) (*athena.TagResourceOutput, error) {
	panic("unimplemented")
}

// TagResourceRequest implements athenaiface.AthenaAPI.
func (*mockAthena) TagResourceRequest(*athena.TagResourceInput) (*request.Request, *athena.TagResourceOutput) {
	panic("unimplemented")
}

// TagResourceWithContext implements athenaiface.AthenaAPI.
func (*mockAthena) TagResourceWithContext(context.Context, *athena.TagResourceInput, ...request.Option) (*athena.TagResourceOutput, error) {
	panic("unimplemented")
}

// UntagResource implements athenaiface.AthenaAPI.
func (*mockAthena) UntagResource(*athena.UntagResourceInput) (*athena.UntagResourceOutput, error) {
	panic("unimplemented")
}

// UntagResourceRequest implements athenaiface.AthenaAPI.
func (*mockAthena) UntagResourceRequest(*athena.UntagResourceInput) (*request.Request, *athena.UntagResourceOutput) {
	panic("unimplemented")
}

// UntagResourceWithContext implements athenaiface.AthenaAPI.
func (*mockAthena) UntagResourceWithContext(context.Context, *athena.UntagResourceInput, ...request.Option) (*athena.UntagResourceOutput, error) {
	panic("unimplemented")
}

// UpdateWorkGroup implements athenaiface.AthenaAPI.
func (*mockAthena) UpdateWorkGroup(*athena.UpdateWorkGroupInput) (*athena.UpdateWorkGroupOutput, error) {
	panic("unimplemented")
}

// UpdateWorkGroupRequest implements athenaiface.AthenaAPI.
func (*mockAthena) UpdateWorkGroupRequest(*athena.UpdateWorkGroupInput) (*request.Request, *athena.UpdateWorkGroupOutput) {
	panic("unimplemented")
}

// UpdateWorkGroupWithContext implements athenaiface.AthenaAPI.
func (*mockAthena) UpdateWorkGroupWithContext(context.Context, *athena.UpdateWorkGroupInput, ...request.Option) (*athena.UpdateWorkGroupOutput, error) {
	panic("unimplemented")
}

func NewMockAthena() athenaiface.AthenaAPI {
	return &mockAthena{}
}
