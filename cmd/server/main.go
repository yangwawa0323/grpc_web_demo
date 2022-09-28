package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/yangwawa0323/grpc_web_demo/pb"
	"github.com/yangwawa0323/grpc_web_demo/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	port := flag.Int("port", 8090, "the server port")
	flag.Parse()

	log.Printf("Start server on port %d\n", *port)

	grpcServer := grpc.NewServer()

	echoServer := service.NewEchoServer()

	pb.RegisterEchoServiceServer(grpcServer, echoServer)

	reflection.Register(grpcServer)

	address := fmt.Sprintf("0.0.0.0:%d", *port)

	listener, err := net.Listen("tcp", address)

	if err != nil {
		log.Fatalf("Cannot listen on port %s\n", address)
	}

	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("Cannot start grpc server.")
	}
}
