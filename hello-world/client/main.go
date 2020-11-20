package main

import (
	"context"
	proto "learn-microservice/hello-world/server/proto"
	"log"

	micro "github.com/micro/go-micro/v2"
)

func main() {
	service := micro.NewService(micro.Name("hello-world"))
	service.Init()

	client := proto.NewGreetService("hello-world", service.Client())
	rsp, err := client.Greet(context.Background(), &proto.GreetRequest{
		Greet: &proto.Greeting{
			FirstName: "Hello",
			LastName: "World",
		},
	})
	if err != nil {
		log.Println(err)
	}
	
	log.Println("Here it comes", rsp.Resutl)
}
