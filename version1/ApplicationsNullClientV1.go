package version1

import (
	"context"

	"github.com/pip-services3-gox/pip-services3-commons-gox/data"
)

type ApplicationsNullClientV1 struct {
}

func NewApplicationsNullClientV1() *ApplicationsNullClientV1 {
	return &ApplicationsNullClientV1{}
}

func (c *ApplicationsNullClientV1) GetApplications(ctx context.Context, correlationId string, filter *data.FilterParams, paging *data.PagingParams) (result data.DataPage[*ApplicationV1], err error) {
	return *data.NewEmptyDataPage[*ApplicationV1](), nil
}

func (c *ApplicationsNullClientV1) GetApplicationById(ctx context.Context, correlationId string, id string) (result *ApplicationV1, err error) {
	return nil, nil
}

func (c *ApplicationsNullClientV1) CreateApplication(ctx context.Context, correlationId string, application *ApplicationV1) (result *ApplicationV1, err error) {
	return application, nil
}

func (c *ApplicationsNullClientV1) UpdateApplication(ctx context.Context, correlationId string, application *ApplicationV1) (result *ApplicationV1, err error) {
	return nil, nil
}

func (c *ApplicationsNullClientV1) DeleteApplicationById(ctx context.Context, correlationId string, id string) (result *ApplicationV1, err error) {
	return nil, nil
}
