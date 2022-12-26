package version1

import (
	"context"

	"github.com/pip-services-content2/client-applications-go/protos"
	"github.com/pip-services3-gox/pip-services3-commons-gox/data"
	"github.com/pip-services3-gox/pip-services3-grpc-gox/clients"
)

type ApplicationGrpcClientV1 struct {
	*clients.GrpcClient
}

func NewApplicationsGrpcClientV1() *ApplicationGrpcClientV1 {
	return &ApplicationGrpcClientV1{
		GrpcClient: clients.NewGrpcClient("applications_v1.Applications"),
	}
}

func (c *ApplicationGrpcClientV1) GetApplications(ctx context.Context, correlationId string, filter *data.FilterParams,
	paging *data.PagingParams) (result data.DataPage[*ApplicationV1], err error) {
	timing := c.Instrument(ctx, correlationId, "applications_v1.get_applications")
	defer timing.EndTiming(ctx, err)

	req := &protos.ApplicationPageRequest{
		CorrelationId: correlationId,
	}
	if filter != nil {
		req.Filter = filter.Value()
	}
	if paging != nil {
		req.Paging = &protos.PagingParams{
			Skip:  paging.GetSkip(0),
			Take:  (int32)(paging.GetTake(100)),
			Total: paging.Total,
		}
	}

	reply := new(protos.ApplicationPageReply)
	err = c.CallWithContext(ctx, "get_applications", correlationId, req, reply)
	if err != nil {
		return *data.NewEmptyDataPage[*ApplicationV1](), err
	}

	if reply.Error != nil {
		err = toError(reply.Error)
		return *data.NewEmptyDataPage[*ApplicationV1](), err
	}

	result = toApplicationPage(reply.Page)

	return result, nil
}

func (c *ApplicationGrpcClientV1) GetApplicationById(ctx context.Context, correlationId string, id string) (result *ApplicationV1, err error) {
	timing := c.Instrument(ctx, correlationId, "applications_v1.get_application_by_id")
	defer timing.EndTiming(ctx, err)

	req := &protos.ApplicationIdRequest{
		CorrelationId: correlationId,
		ApplicationId: id,
	}

	reply := new(protos.ApplicationObjectReply)
	err = c.CallWithContext(ctx, "get_application_by_id", correlationId, req, reply)
	if err != nil {
		return nil, err
	}

	if reply.Error != nil {
		err = toError(reply.Error)
		return nil, err
	}

	result = toApplication(reply.Application)

	return result, nil
}

func (c *ApplicationGrpcClientV1) CreateApplication(ctx context.Context, correlationId string, application *ApplicationV1) (result *ApplicationV1, err error) {
	timing := c.Instrument(ctx, correlationId, "applications_v1.create_application")
	defer timing.EndTiming(ctx, err)

	req := &protos.ApplicationObjectRequest{
		CorrelationId: correlationId,
		Application:   fromApplication(application),
	}

	reply := new(protos.ApplicationObjectReply)
	err = c.CallWithContext(ctx, "create_application", correlationId, req, reply)
	if err != nil {
		return nil, err
	}

	if reply.Error != nil {
		err = toError(reply.Error)
		return nil, err
	}

	result = toApplication(reply.Application)

	return result, nil
}

func (c *ApplicationGrpcClientV1) UpdateApplication(ctx context.Context, correlationId string, application *ApplicationV1) (result *ApplicationV1, err error) {
	timing := c.Instrument(ctx, correlationId, "applications_v1.update_application")
	defer timing.EndTiming(ctx, err)

	req := &protos.ApplicationObjectRequest{
		CorrelationId: correlationId,
		Application:   fromApplication(application),
	}

	reply := new(protos.ApplicationObjectReply)
	err = c.CallWithContext(ctx, "update_application", correlationId, req, reply)
	if err != nil {
		return nil, err
	}

	if reply.Error != nil {
		err = toError(reply.Error)
		return nil, err
	}

	result = toApplication(reply.Application)

	return result, nil
}

func (c *ApplicationGrpcClientV1) DeleteApplicationById(ctx context.Context, correlationId string, id string) (result *ApplicationV1, err error) {
	timing := c.Instrument(ctx, correlationId, "applications_v1.delete_application_by_id")
	defer timing.EndTiming(ctx, err)

	req := &protos.ApplicationIdRequest{
		CorrelationId: correlationId,
		ApplicationId: id,
	}

	reply := new(protos.ApplicationObjectReply)
	err = c.CallWithContext(ctx, "delete_application_by_id", correlationId, req, reply)
	if err != nil {
		return nil, err
	}

	if reply.Error != nil {
		err = toError(reply.Error)
		return nil, err
	}

	result = toApplication(reply.Application)

	return result, nil
}
