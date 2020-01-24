package server

import (
	"context"

	"github.com/micro/go-micro"
	"github.com/owncloud/ocis-pkg/service/grpc"
	"github.com/refs/example-extension/pkg/proto"
	svc "github.com/refs/example-extension/pkg/service"
)

// NewServer returns a go-micro server
func NewServer(c context.Context) micro.Service {
	service := grpc.NewService(
		grpc.Context(c),
		grpc.Namespace("com.owncloud.api"),
		grpc.Name("greeter"),
		grpc.Address("localhost:10001"),
	)

	service.Init()

	proto.RegisterGreeterServiceHandler(service.Server(), svc.Service{})

	return service
}
