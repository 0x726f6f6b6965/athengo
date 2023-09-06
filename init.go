package athengo

import (
	"athengo/config"
	"context"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/athena"
	"github.com/aws/aws-sdk-go/service/athena/athenaiface"
)

type athenaClient struct {
	athena         athenaiface.AthenaAPI
	database       string
	outputLocation string
	pollFrequency  time.Duration
}

type AthenaClient interface {
	ExecQuery(ctx context.Context, query string) (Rows, error)
	GetAthenaApi() athenaiface.AthenaAPI
}

// This is aws db service
// Require AWS credentials.
// The simplest way to provide them is via AWS_ACCESS_KEY_ID
// and AWS_SECRET_ACCESS_KEY environment variables.
// Doc: https://docs.aws.amazon.com/sdk-for-java/v1/developer-guide/credentials.html#credentials-default
func NewAthenaClient(cfg *config.AthenaConfig) AthenaClient {
	awsCfg := &aws.Config{}
	awsCfg.WithRegion(cfg.Region)
	conn := session.Must(session.NewSession(awsCfg))
	return NewAthenaClientWithSession(conn, cfg)
}

func NewAthenaClientWithSession(conn *session.Session, cfg *config.AthenaConfig) AthenaClient {
	api := athena.New(conn)
	client := &athenaClient{
		athena:         api,
		database:       cfg.DbName,
		outputLocation: cfg.OutputLocation,
		pollFrequency:  cfg.PollFrequency,
	}
	return client
}

func NewAthenaClientWithAPI(api athenaiface.AthenaAPI, cfg *config.AthenaConfig) AthenaClient {
	client := &athenaClient{
		athena:         api,
		database:       cfg.DbName,
		outputLocation: cfg.OutputLocation,
		pollFrequency:  cfg.PollFrequency,
	}
	return client
}

func (c *athenaClient) GetAthenaApi() athenaiface.AthenaAPI {
	return c.athena
}
