package test_version1

import (
	"testing"

	"github.com/pip-services-content2/client-applications-go/version1"
)

type applicationsMockClientV1Test struct {
	client  *version1.ApplicationsMockClientV1
	fixture *ApplicationsClientFixtureV1
}

func newApplicationsMockClientV1Test() *applicationsMockClientV1Test {
	return &applicationsMockClientV1Test{}
}

func (c *applicationsMockClientV1Test) setup(t *testing.T) *ApplicationsClientFixtureV1 {
	c.client = version1.NewApplicationsMockClientV1()
	c.fixture = NewApplicationsClientFixtureV1(c.client)

	return c.fixture
}

func (c *applicationsMockClientV1Test) teardown(t *testing.T) {
	c.client = nil
}

func TestMockOperations(t *testing.T) {
	c := newApplicationsMockClientV1Test()
	fixture := c.setup(t)
	defer c.teardown(t)

	fixture.TestCrudOperations(t)
}
