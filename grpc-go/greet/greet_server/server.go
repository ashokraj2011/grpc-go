package main

import (
	"fmt"
	"log"
	"net"

	"github.com/ashokraj1978/grpc-go/grpc-go/greet/greetpd"
	"github.com/ashokraj1978/grpc-go/tree/main/grpc-go/greet/greetpd"

	grpc "google.golang.org/grpc"
)

type server struct{}

func main() {
	fmt.Println("hai")
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("cannot listn", err)
	}
	s := grpc.NewServer()
	greetpd.RegisterGreetServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to server %v ", err)
	}
}
