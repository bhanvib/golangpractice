package main

import (
	"context"
	"fmt"

	"project.com/proto/greetpb"
	"google.golang.org/grpc"
)
func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	defer conn.Close()
	if err != nil {
		fmt.Println(err)
	}
	c := greetpb.NewGreetServiceClient(conn)
	fmt.Println("client is created")
	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Stephen",
			LastName:  "Mark",
		},
	}
	res, err := c.Greet(context.Background(), req)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res.GetResult())
}