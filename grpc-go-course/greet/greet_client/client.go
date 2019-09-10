package main

import (
	"context"
	"fmt"
	"log"

	"github.com/eltneg/grpc-go-course/greet/greetpb"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer conn.Close()
	c := greetpb.NewGreetServiceClient(conn)
	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "elt",
			LastName:  "neg",
		},
	}
	res, err := c.Greet(context.Background(), req)
	if err != nil {
		fmt.Print("error", err)
	}
	log.Printf("Response from greet: %v", res.Result)
}
