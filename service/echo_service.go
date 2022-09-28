package service

import (
	"context"
	"fmt"
	"os"

	"github.com/yangwawa0323/grpc_web_demo/pb"
)

type EchoServer struct {
	pb.UnimplementedEchoServiceServer
}

func NewEchoServer() *EchoServer {
	return &EchoServer{}
}

func (es *EchoServer) Echo(ctx context.Context,
	req *pb.EchoRequest) (*pb.EchoResponse, error) {

	fmt.Fprintf(os.Stdout, "Recieved the message : %s", req.Message)
	return &pb.EchoResponse{
		Message: "Hello " + req.Message,
	}, nil
}
