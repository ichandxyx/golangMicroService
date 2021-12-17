package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ichandxyx/micro/greet/greetpb"
	"google.golang.org/grpc"
)
func main(){
	fmt.Println("hello I am client")
	conn, err:= grpc.Dial("localhost:50051",grpc.WithInsecure())
	if err!=nil {
		log.Fatal("not able to dila connection : %v", err)
	}

	defer conn.Close()
	c:= greetpb.NewGreetServiceClient(conn)

	doUnary(c)
	fmt.Printf("created client: %v", c)
}

func doUnary(c greetpb.GreetServiceClient){
	fmt.Println("starting to do a Unary RPC...")
req:= greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "chandan",
			LastName: "kumar",
		},
	}
	res,err:=c.Greet(context.Background(), &req)
	if err!=nil{
		log.Fatalf("error while calling Greet RPC %v", err)
	}
	log.Printf("Response from greet :%v", res.Result)
}