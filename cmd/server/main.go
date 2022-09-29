package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	pb_echo "github.com/yangwawa0323/grpc_web_demo/pb/echo/v1"
	pb_user "github.com/yangwawa0323/grpc_web_demo/pb/user/v1"
	"github.com/yangwawa0323/grpc_web_demo/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	port := flag.Int("port", 8090, "the server port")
	flag.Parse()

	log.Printf("Start server on port %d\n", *port)

	grpcServer := grpc.NewServer()

	// Create service servers
	echoServer := service.NewEchoServer()

	userServer := service.NewUserServiceServer()

	// Register service servers to grpc server
	pb_echo.RegisterEchoServiceServer(grpcServer, echoServer)

	pb_user.RegisterUserSearchServiceServer(grpcServer, userServer)

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
