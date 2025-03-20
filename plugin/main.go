package main

import (
	"context"
	"log"
	"net"

	pb "github.com/kairveeehh/grpc-plugin-poc/proto"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedVMPluginServer
}

func (s *server) Start(ctx context.Context, in *pb.Empty) (*pb.Response, error) {
	log.Println("Plugin: Start called")
	return &pb.Response{Message: "VM Started"}, nil
}

func (s *server) Stop(ctx context.Context, in *pb.Empty) (*pb.Response, error) {
	log.Println("Plugin: Stop called")
	return &pb.Response{Message: "VM Stopped"}, nil
}

func (s *server) Info(ctx context.Context, in *pb.Empty) (*pb.Response, error) {
	log.Println("Plugin: Info called")
	return &pb.Response{Message: "VM Info: Running"}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterVMPluginServer(s, &server{})
	log.Println("Plugin gRPC server running on port 50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
