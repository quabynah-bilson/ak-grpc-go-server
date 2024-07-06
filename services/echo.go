package services

import (
	"context"
	"fmt"
	pb "github.com/quabynah-bilson/grpc-go/protos"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type echoServiceImpl struct {
	pb.UnimplementedEchoServiceServer
}

func NewEchoService() pb.EchoServiceServer {
	return &echoServiceImpl{}
}

func (e *echoServiceImpl) EchoMessage(_ context.Context, req *pb.EchoRequest) (*pb.EchoResponse, error) {
	// validate the request
	if len(req.GetName()) == 0 {
		return nil, status.Error(codes.InvalidArgument, "name cannot be empty")
	}
	
	// create response with greeting
	response := &pb.EchoResponse{Greeting: fmt.Sprintf("Hello %s! Welcome to your first grpc server", req.GetName())}
	
	// return response
	return response, nil
}
