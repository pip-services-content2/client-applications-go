package test_version1

import (
	"context"
	"os"
	"testing"

	"github.com/pip-services-content2/client-applications-go/version1"
	"github.com/pip-services3-gox/pip-services3-commons-gox/config"
)

type applicationsGrpcClientV1Test struct {
	client  *version1.ApplicationGrpcClientV1
	fixture *ApplicationsClientFixtureV1
}

func newApplicationsGrpcClientV1Test() *applicationsGrpcClientV1Test {
	return &applicationsGrpcClientV1Test{}
}

func (c *applicationsGrpcClientV1Test) setup(t *testing.T) *ApplicationsClientFixtureV1 {
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

	c.client = version1.NewApplicationsGrpcClientV1()
	c.client.Configure(context.Background(), httpConfig)
	c.client.Open(context.Background(), "")

	c.fixture = NewApplicationsClientFixtureV1(c.client)

	return c.fixture
}

func (c *applicationsGrpcClientV1Test) teardown(t *testing.T) {
	c.client.Close(context.Background(), "")
}

func TestGrpcCrudOperations(t *testing.T) {
	c := newApplicationsGrpcClientV1Test()
	fixture := c.setup(t)
	defer c.teardown(t)

	fixture.TestCrudOperations(t)
}
