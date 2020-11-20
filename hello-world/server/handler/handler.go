package handler

import (
	"context"
	hello "learn-microservice/hello-world/server/proto"
)

type Helloworld struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Helloworld) Greet(ctx context.Context, req *hello.GreetRequest, rsp *hello.GreetResponse) error {
	rsp.Resutl = "Hello " + req.Greet.FirstName + req.Greet.LastName
	return nil
}