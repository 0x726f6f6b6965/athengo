package athengo

import (
	"athengo/config"
	"athengo/mock"
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go/service/athena"
	"github.com/stretchr/testify/assert"
)

var Client AthenaClient

func setup() {
	// // please enter the access key
	// os.Setenv("AWS_ACCESS_KEY_ID", "")
	// // please enter the secret access key
	// os.Setenv("AWS_SECRET_ACCESS_KEY", "")

	// Client = NewAthenaClient(config.NewConfig("region", "dbName", "outputLocation", time.Millisecond*100))
	Client = NewAthenaClientWithAPI(mock.NewMockAthena(), config.NewConfig("region", "dbName", "outputLocation", time.Millisecond*100))
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	os.Exit(code)
}

func TestMockGetResult(t *testing.T) {
	resp, err := Client.ExecQuery(context.Background(),
		`SELECT id, name, create_time, update_time FROM mock`)
	assert.Equal(t, true, err == nil, fmt.Sprintf("the err is not nil, err: %v", err))
	if err != nil {
		return
	}
	result := []*MockData{}
	err = resp.GetResults(&result)

	assert.Equal(t, true, err == nil, fmt.Sprintf("the err is not nil, err: %v", err))
	if err != nil {
		return
	}
	assert.Equal(t, true, len(result) == 3, "the length of result is wrong")
}

func TestMockFailQuery(t *testing.T) {
	ctx := context.WithValue(context.Background(), mock.STATUS, athena.QueryExecutionStateFailed)
	_, err := Client.ExecQuery(ctx,
		`SELECT id, name, create_time, update_time FROM mock`)
	assert.Equal(t, true, err != nil, "the err is nil")
}

func TestMockCancelledQuery(t *testing.T) {
	ctx := context.WithValue(context.Background(), mock.STATUS, athena.QueryExecutionStateCancelled)
	_, err := Client.ExecQuery(ctx,
		`SELECT id, name, create_time, update_time FROM mock`)
	assert.Equal(t, true, err != nil, "the err is nil")
}

type MockData struct {
	Id         uint64    `athena:"id"`
	Name       string    `athena:"name"`
	CreateTime time.Time `athena:"create_time"`
	UpdateTime time.Time `athena:"update_time"`
}
