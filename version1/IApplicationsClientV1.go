package version1

import (
	"context"

	"github.com/pip-services3-gox/pip-services3-commons-gox/data"
)

type IApplicationsClientV1 interface {
	GetApplications(ctx context.Context, correlationId string, filter *data.FilterParams,
		paging *data.PagingParams) (result data.DataPage[*ApplicationV1], err error)

	GetApplicationById(ctx context.Context, correlationId string, id string) (result *ApplicationV1, err error)

	CreateApplication(ctx context.Context, correlationId string, application *ApplicationV1) (result *ApplicationV1, err error)

	UpdateApplication(ctx context.Context, correlationId string, application *ApplicationV1) (result *ApplicationV1, err error)

	DeleteApplicationById(ctx context.Context, correlationId string, id string) (result *ApplicationV1, err error)
}
