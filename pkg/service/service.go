package service

import (
	"context"

	"github.com/refs/example-extension/pkg/proto"
)

// Service GreeterServiceHandler
type Service struct{}

// Greet implements the GreeterServiceHandler interface
func (s Service) Greet(c context.Context, req *proto.GreetRequest, res *proto.GreetResponse) error {
	// insert business logic here
	res.Resp = "whatever"
	return nil
}
