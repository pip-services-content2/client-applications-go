package build

import (
	clients1 "github.com/pip-services-content2/client-applications-go/version1"
	cref "github.com/pip-services3-gox/pip-services3-commons-gox/refer"
	cbuild "github.com/pip-services3-gox/pip-services3-components-gox/build"
)

type EmailClientFactory struct {
	*cbuild.Factory
}

func NewEmailClientFactory() *EmailClientFactory {
	c := &EmailClientFactory{
		Factory: cbuild.NewFactory(),
	}

	nullClientDescriptor := cref.NewDescriptor("service-applications", "client", "null", "*", "1.0")
	cmdHttpClientDescriptor := cref.NewDescriptor("service-applications", "client", "commandable-http", "*", "1.0")
	grpcClientDescriptor := cref.NewDescriptor("service-applications", "client", "grpc", "*", "1.0")
	cmdGrpcClientDescriptor := cref.NewDescriptor("service-applications", "client", "commandable-grpc", "*", "1.0")

	c.RegisterType(nullClientDescriptor, clients1.NewApplicationsNullClientV1)
	c.RegisterType(cmdHttpClientDescriptor, clients1.NewApplicationsCommandableHttpClientV1)
	c.RegisterType(grpcClientDescriptor, clients1.NewApplicationsGrpcClientV1)
	c.RegisterType(cmdGrpcClientDescriptor, clients1.NewApplicationsCommandableGrpcClientV1)

	return c
}
