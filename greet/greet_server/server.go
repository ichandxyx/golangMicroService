package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/ichandxyx/micro/greet/greetpb"
	"google.golang.org/grpc"
)

type server struct{}

func (*server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	fmt.Println("into the Greet with request : ", req)
	firstName := req.GetGreeting().GetFirstName()
	result := "hello " + firstName
	res := &greetpb.GreetResponse{
		Result: &greetpb.Greeting{
			FirstName: result,
		},
	}
	return res, nil
}

func (*server) Add(ctx context.Context, req *greetpb.AddRequest) (*greetpb.AddResponse, error) {
	fmt.Println("into the Add With request : ", req)
	num1 := req.GetNumbers().GetFirstNumber()
	num2 := req.GetNumbers().GetSecondNumber()
	result := num1 + num2
	res := &greetpb.AddResponse{
		Result: &greetpb.Add{
			FirstNumber: result,
		},
	}
	return res, nil
}

func main() {
	fmt.Println("Hello World")
	list, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, &server{})
	greetpb.RegisterAddServiceServer(s, &server{})
	if err := s.Serve((list)); err != nil {
		log.Fatalf("failed to serve : %v", err)
	}
}
