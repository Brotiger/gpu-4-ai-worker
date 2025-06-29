package main

import (
	"log"
	"net"

	"gpu-4-ai-worker/internal/config"
	"gpu-4-ai-worker/internal/handler"

	"github.com/Brotiger/gpu-4-ai-worker/proto"
	"google.golang.org/grpc"
)

func main() {
	cfg := config.Load()
	lis, err := net.Listen("tcp", cfg.GRPCAddr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	proto.RegisterWorkerServer(grpcServer, handler.NewWorkerHandler(cfg))
	log.Println("gRPC server listening on", cfg.GRPCAddr)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
