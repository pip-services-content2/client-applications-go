package test_version1

import (
	"context"
	"os"
	"testing"

	"github.com/pip-services-content2/client-applications-go/version1"
	"github.com/pip-services3-gox/pip-services3-commons-gox/config"
)

type applicationsCommandableGrpcClientV1Test struct {
	client  *version1.ApplicationsCommandableGrpcClientV1
	fixture *ApplicationsClientFixtureV1
}

func newApplicationsCommandableGrpcClientV1Test() *applicationsCommandableGrpcClientV1Test {
	return &applicationsCommandableGrpcClientV1Test{}
}

func (c *applicationsCommandableGrpcClientV1Test) setup(t *testing.T) *ApplicationsClientFixtureV1 {
	var GRPC_HOST = os.Getenv("GRPC_HOST")
	if GRPC_HOST == "" {
		GRPC_HOST = "localhost"
	}
	var GRPC_PORT = os.Getenv("GRPC_PORT")
	if GRPC_PORT == "" {
		GRPC_PORT = "8090"
	}

	var httpConfig = config.NewConfigParamsFromTuples(
		"connection.protocol", "http",
		"connection.host", GRPC_HOST,
		"connection.port", GRPC_PORT,
	)

	c.client = version1.NewApplicationsCommandableGrpcClientV1()
	c.client.Configure(context.Background(), httpConfig)
	c.client.Open(context.Background(), "")

	c.fixture = NewApplicationsClientFixtureV1(c.client)

	return c.fixture
}

func (c *applicationsCommandableGrpcClientV1Test) teardown(t *testing.T) {
	c.client.Close(context.Background(), "")
}

func TestCommandableGrpcCrudOperations(t *testing.T) {
	c := newApplicationsCommandableGrpcClientV1Test()
	fixture := c.setup(t)
	defer c.teardown(t)

	fixture.TestCrudOperations(t)
}
