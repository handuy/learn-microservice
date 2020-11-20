// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: hello.proto

package proto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/micro/go-micro/v2/api"
	client "github.com/micro/go-micro/v2/client"
	server "github.com/micro/go-micro/v2/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for GreetService service

func NewGreetServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for GreetService service

type GreetService interface {
	Greet(ctx context.Context, in *GreetRequest, opts ...client.CallOption) (*GreetResponse, error)
}

type greetService struct {
	c    client.Client
	name string
}

func NewGreetService(name string, c client.Client) GreetService {
	return &greetService{
		c:    c,
		name: name,
	}
}

func (c *greetService) Greet(ctx context.Context, in *GreetRequest, opts ...client.CallOption) (*GreetResponse, error) {
	req := c.c.NewRequest(c.name, "GreetService.Greet", in)
	out := new(GreetResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for GreetService service

type GreetServiceHandler interface {
	Greet(context.Context, *GreetRequest, *GreetResponse) error
}

func RegisterGreetServiceHandler(s server.Server, hdlr GreetServiceHandler, opts ...server.HandlerOption) error {
	type greetService interface {
		Greet(ctx context.Context, in *GreetRequest, out *GreetResponse) error
	}
	type GreetService struct {
		greetService
	}
	h := &greetServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&GreetService{h}, opts...))
}

type greetServiceHandler struct {
	GreetServiceHandler
}

func (h *greetServiceHandler) Greet(ctx context.Context, in *GreetRequest, out *GreetResponse) error {
	return h.GreetServiceHandler.Greet(ctx, in, out)
}
