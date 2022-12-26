package test_version1

import (
	"context"
	"testing"

	"github.com/pip-services-content2/client-applications-go/version1"
	"github.com/stretchr/testify/assert"
)

type ApplicationsClientFixtureV1 struct {
	Client version1.IApplicationsClientV1
}

var APPLICATION1 = &version1.ApplicationV1{
	Id:         "1",
	Name:       map[string]string{"en": "App 1"},
	Product:    "Product 1",
	Copyrights: "PipDevs 2018",
	MinVer:     0,
	MaxVer:     9999,
}
var APPLICATION2 = &version1.ApplicationV1{
	Id:         "2",
	Name:       map[string]string{"en": "App 2"},
	Product:    "Product 1",
	Copyrights: "PipDevs 2018",
	MinVer:     0,
	MaxVer:     9999,
}

func NewApplicationsClientFixtureV1(client version1.IApplicationsClientV1) *ApplicationsClientFixtureV1 {
	return &ApplicationsClientFixtureV1{
		Client: client,
	}
}

func (c *ApplicationsClientFixtureV1) clear() {
	page, _ := c.Client.GetApplications(context.Background(), "", nil, nil)

	for _, v := range page.Data {
		application := v
		c.Client.DeleteApplicationById(context.Background(), "", application.Id)
	}
}

func (c *ApplicationsClientFixtureV1) TestCrudOperations(t *testing.T) {
	c.clear()
	defer c.clear()

	// Create one application
	application, err := c.Client.CreateApplication(context.Background(), "", APPLICATION1)
	assert.Nil(t, err)

	assert.NotNil(t, application)
	assert.Equal(t, application.Product, APPLICATION1.Product)
	assert.Equal(t, application.Copyrights, APPLICATION1.Copyrights)

	application1 := application

	// Create another application
	application, err = c.Client.CreateApplication(context.Background(), "", APPLICATION2)
	assert.Nil(t, err)

	assert.NotNil(t, application)
	assert.Equal(t, application.Product, APPLICATION2.Product)
	assert.Equal(t, application.Copyrights, APPLICATION2.Copyrights)

	//application2 := application

	// Get all applications
	page, err1 := c.Client.GetApplications(context.Background(), "", nil, nil)
	assert.Nil(t, err1)

	assert.NotNil(t, page)
	assert.True(t, len(page.Data) >= 2)

	// Update the application
	application1.Name = map[string]string{"en": "Updated Name 1"}
	application, err = c.Client.UpdateApplication(context.Background(), "", application1)
	assert.Nil(t, err)

	assert.NotNil(t, application)
	assert.Equal(t, application.Name["en"], "Updated Name 1")
	assert.Equal(t, application.Id, application1.Id)

	application1 = application

	// Delete application
	application, err = c.Client.DeleteApplicationById(context.Background(), "", application1.Id)
	assert.Nil(t, err)

	// Try to get deleted application
	application, err = c.Client.GetApplicationById(context.Background(), "", application1.Id)
	assert.Nil(t, err)

	assert.Nil(t, application)
}
