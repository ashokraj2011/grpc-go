package main

import (
	"fmt"
	"log"
	"net"

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
	greetpd.
}
