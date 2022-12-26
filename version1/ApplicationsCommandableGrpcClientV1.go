package version1

import (
	"context"

	cconf "github.com/pip-services3-gox/pip-services3-commons-gox/config"
	cdata "github.com/pip-services3-gox/pip-services3-commons-gox/data"
	"github.com/pip-services3-gox/pip-services3-grpc-gox/clients"
)

type ApplicationsCommandableGrpcClientV1 struct {
	*clients.CommandableGrpcClient
}

func NewApplicationsCommandableGrpcClientV1() *ApplicationsCommandableGrpcClientV1 {
	return NewApplicationsCommandableGrpcClientV1WithConfig(nil)
}

func NewApplicationsCommandableGrpcClientV1WithConfig(config *cconf.ConfigParams) *ApplicationsCommandableGrpcClientV1 {
	c := &ApplicationsCommandableGrpcClientV1{
		CommandableGrpcClient: clients.NewCommandableGrpcClient("v1/applications"),
	}

	if config != nil {
		c.Configure(context.Background(), config)
	}

	return c
}

func (c *ApplicationsCommandableGrpcClientV1) GetApplications(ctx context.Context, correlationId string, filter *cdata.FilterParams,
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

func (c *ApplicationsCommandableGrpcClientV1) GetApplicationById(ctx context.Context, correlationId string, id string) (result *ApplicationV1, err error) {
	res, err := c.CallCommand(ctx, "get_application_by_id", correlationId, cdata.NewAnyValueMapFromTuples(
		"application_id", id,
	))

	if err != nil {
		return nil, err
	}

	return clients.HandleHttpResponse[*ApplicationV1](res, correlationId)
}

func (c *ApplicationsCommandableGrpcClientV1) CreateApplication(ctx context.Context, correlationId string, application *ApplicationV1) (result *ApplicationV1, err error) {
	res, err := c.CallCommand(ctx, "create_application", correlationId, cdata.NewAnyValueMapFromTuples(
		"application", application,
	))

	if err != nil {
		return nil, err
	}

	return clients.HandleHttpResponse[*ApplicationV1](res, correlationId)
}

func (c *ApplicationsCommandableGrpcClientV1) UpdateApplication(ctx context.Context, correlationId string, application *ApplicationV1) (result *ApplicationV1, err error) {
	res, err := c.CallCommand(ctx, "update_application", correlationId, cdata.NewAnyValueMapFromTuples(
		"application", application,
	))

	if err != nil {
		return nil, err
	}

	return clients.HandleHttpResponse[*ApplicationV1](res, correlationId)
}

func (c *ApplicationsCommandableGrpcClientV1) DeleteApplicationById(ctx context.Context, correlationId string, id string) (result *ApplicationV1, err error) {
	res, err := c.CallCommand(ctx, "delete_application_by_id", correlationId, cdata.NewAnyValueMapFromTuples(
		"application_id", id,
	))

	if err != nil {
		return nil, err
	}

	return clients.HandleHttpResponse[*ApplicationV1](res, correlationId)
}
