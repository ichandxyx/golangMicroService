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

func (*server) Greet(ctx context.Context, req *greetpb.GreetRequest) ( *greetpb.GreetResponse, error){
	fmt.Println("into the Greet with request : %v", req)
	firstName:= req.GetGreeting().GetFirstName()
	result:= "hello "+ firstName
	res:= &greetpb.GreetResponse{
		Result: &greetpb.Greeting{
			FirstName: result,
		},
	}
	return res,nil
}

func main(){
	fmt.Println("Hello World")
	list,err:= net.Listen("tcp", "0.0.0.0:50051")
	if err!=nil{
		log.Fatal("failed to listen: %v", err)
	}
	s:= grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, &server{})
	if err:=s.Serve((list)); err !=nil {
		log.Fatal("failed to serve : %v",err)
	}
}