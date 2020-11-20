package main

import (
	"learn-microservice/hello-world/server/handler"
	"learn-microservice/hello-world/server/proto"
	"log"

	micro "github.com/micro/go-micro/v2"
)

func main() {
	// Create a new service. Optionally include some options here.
	service := micro.NewService(

		// This name must match the package name given in your protobuf definition
		micro.Name("hello-world"),
	)

	// Init will parse the command line flags.
	service.Init()

	// Register service
	if err := proto.RegisterGreetServiceHandler(service.Server(), &handler.Helloworld{}); err != nil {
		log.Panic(err)
	}

	// Run the server
	if err := service.Run(); err != nil {
		log.Panic(err)
	}
}
