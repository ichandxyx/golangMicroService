package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ichandxyx/micro/greet/greetpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("hello I am client")
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("not able to dial connection : %v", err)
	}
	defer conn.Close()

	c := greetpb.NewGreetServiceClient(conn)
	s := greetpb.NewAddServiceClient(conn)

	var option int

	fmt.Println("Enter\n1 For Greet service \n2 For Add service")
	fmt.Scanln(&option)

	switch option {
	case 1:
		doUnary(c)
	case 2:
		doSum(s)
	}

}

func doUnary(c greetpb.GreetServiceClient) {
	fmt.Println("starting to do a Unary RPC...")
	req := greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "chandan",
			LastName:  "kumar",
		},
	}
	res, err := c.Greet(context.Background(), &req)
	if err != nil {
		log.Fatalf("error while calling Greet RPC %v", err)
	}
	log.Printf("Response from greet :%v", res.Result)
}

func doSum(c greetpb.AddServiceClient) {
	fmt.Println("starting to do a Add RPC...")
	var num1, num2 int32
	
	fmt.Println("Enter first number")
	fmt.Scanln(&num1)

	fmt.Println("Enter second number")
	fmt.Scanln(&num2)

	req := greetpb.AddRequest{
		Numbers: &greetpb.Add{
			FirstNumber:  num1,
			SecondNumber: num2,
		},
	}
	res, err := c.Add(context.Background(), &req)
	if err != nil {
		log.Fatalf("Error while calling Add RPC: %v", err)
	}
	log.Printf("The sum Calculated is : %v ", res.Result.FirstNumber)
}
