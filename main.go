package main

import (
	"context"
	"log"
	"os/exec"
	"time"

	pb "github.com/kairveeehh/grpc-plugin-poc/proto"
	"google.golang.org/grpc"
)

func main() {
	cmd := exec.Command("./plugin/vmplugin")
	err := cmd.Start()
	if err != nil {
		log.Fatalf("Failed to start plugin: %v", err)
	}
	defer cmd.Process.Kill()

	time.Sleep(1 * time.Second) // Wait for plugin to boot

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to plugin: %v", err)
	}
	defer conn.Close()

	client := pb.NewVMPluginClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, _ := client.Start(ctx, &pb.Empty{})
	log.Println("Start Response:", res.Message)

	res, _ = client.Info(ctx, &pb.Empty{})
	log.Println("Info Response:", res.Message)

	res, _ = client.Stop(ctx, &pb.Empty{})
	log.Println("Stop Response:", res.Message)
}
