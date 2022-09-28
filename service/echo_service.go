package service

import (
	"context"
	"fmt"
	"os"
	"time"

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

	fmt.Fprintf(os.Stdout, "%s: Recieved the message : %s\n", time.Now().Format("2006-01-02 15:04:05"), req.Message)
	return &pb.EchoResponse{
		Message: "Hello " + req.Message,
	}, nil
}
