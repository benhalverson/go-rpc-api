package main

import (
	"context"
	"fmt"
	"go-grpc-api/greet/greetpb"
	"log"

	"google.golang.org/grpc"
)

func main() {
	fmt.Print("Hello I'm the client")
	opts := grpc.WithInsecure()
	cc, err := grpc.Dial("localhost:50051", opts)

	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer cc.Close()

	c := greetpb.NewGreetServiceClient(cc)
	fmt.Printf("created client: %f", c)
	doUnary(c)
}

func doUnary(c greetpb.GreetServiceClient) {
	println("Starting to do a Unary RPC")
	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Ben",
			LastName:  "Halverson",
		},
	}
	res, err := c.Greet(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling greet RPC %v", err)
	}

	log.Printf("Response from Greet RPC %v", res.Result)
}
