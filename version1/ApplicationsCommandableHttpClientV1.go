package version1

import (
	"context"

	cconf "github.com/pip-services3-gox/pip-services3-commons-gox/config"
	cdata "github.com/pip-services3-gox/pip-services3-commons-gox/data"
	"github.com/pip-services3-gox/pip-services3-rpc-gox/clients"
)

type ApplicationsCommandableHttpClientV1 struct {
	*clients.CommandableHttpClient
}

func NewApplicationsCommandableHttpClientV1() *ApplicationsCommandableHttpClientV1 {
	return NewApplicationsCommandableHttpClientV1WithConfig(nil)
}

func NewApplicationsCommandableHttpClientV1WithConfig(config *cconf.ConfigParams) *ApplicationsCommandableHttpClientV1 {
	c := &ApplicationsCommandableHttpClientV1{
		CommandableHttpClient: clients.NewCommandableHttpClient("v1/applications"),
	}

	if config != nil {
		c.Configure(context.Background(), config)
	}

	return c
}

func (c *ApplicationsCommandableHttpClientV1) GetApplications(ctx context.Context, correlationId string, filter *cdata.FilterParams,
	paging *cdata.PagingParams) (result cdata.DataPage[*ApplicationV1], err error) {
	res, err := c.CallCommand(ctx, "get_applications", correlationId, cdata.NewAnyValueMapFromTuples(
		"filter", filter,
		"paging", paging,
	))

	if err != nil {
		return *cdata.NewEmptyDataPage[*ApplicationV1](), err
	}

	return clients.HandleHttpResponse[cdata.DataPage[*ApplicationV1]](res, correlationId)
}

func (c *ApplicationsCommandableHttpClientV1) GetApplicationById(ctx context.Context, correlationId string, id string) (result *ApplicationV1, err error) {
	res, err := c.CallCommand(ctx, "get_application_by_id", correlationId, cdata.NewAnyValueMapFromTuples(
		"application_id", id,
	))

	if err != nil {
		return nil, err
	}

	return clients.HandleHttpResponse[*ApplicationV1](res, correlationId)
}

func (c *ApplicationsCommandableHttpClientV1) CreateApplication(ctx context.Context, correlationId string, application *ApplicationV1) (result *ApplicationV1, err error) {
	res, err := c.CallCommand(ctx, "create_application", correlationId, cdata.NewAnyValueMapFromTuples(
		"application", application,
	))

	if err != nil {
		return nil, err
	}

	return clients.HandleHttpResponse[*ApplicationV1](res, correlationId)
}

func (c *ApplicationsCommandableHttpClientV1) UpdateApplication(ctx context.Context, correlationId string, application *ApplicationV1) (result *ApplicationV1, err error) {
	res, err := c.CallCommand(ctx, "update_application", correlationId, cdata.NewAnyValueMapFromTuples(
		"application", application,
	))

	if err != nil {
		return nil, err
	}

	return clients.HandleHttpResponse[*ApplicationV1](res, correlationId)
}

func (c *ApplicationsCommandableHttpClientV1) DeleteApplicationById(ctx context.Context, correlationId string, id string) (result *ApplicationV1, err error) {
	res, err := c.CallCommand(ctx, "delete_application_by_id", correlationId, cdata.NewAnyValueMapFromTuples(
		"application_id", id,
	))

	if err != nil {
		return nil, err
	}

	return clients.HandleHttpResponse[*ApplicationV1](res, correlationId)
}
