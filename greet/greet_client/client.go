package main

import (
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
}
