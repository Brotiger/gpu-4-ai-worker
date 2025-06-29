package handler

import (
	"context"

	"gpu-4-ai-worker/internal/config"
	"gpu-4-ai-worker/internal/service"

	workerpb "github.com/Brotiger/gpu-4-ai-worker/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type WorkerHandler struct {
	workerpb.UnimplementedWorkerServer
	Cfg           *config.Config
	OllamaService *service.OllamaService
}

func NewWorkerHandler(cfg *config.Config) *WorkerHandler {
	return &WorkerHandler{
		Cfg:           cfg,
		OllamaService: service.NewOllamaService(cfg),
	}
}

func (h *WorkerHandler) Generate(ctx context.Context, req *workerpb.GenerateRequest) (*workerpb.GenerateResponse, error) {
	resp, err := h.OllamaService.Generate(req)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to generate: %v", err)
	}
	return resp, nil
}

func (h *WorkerHandler) Tags(ctx context.Context, req *workerpb.TagsRequest) (*workerpb.TagsResponse, error) {
	resp, err := h.OllamaService.Tags()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get tags: %v", err)
	}
	return resp, nil
}

func (h *WorkerHandler) Show(ctx context.Context, req *workerpb.ShowRequest) (*workerpb.ShowResponse, error) {
	resp, err := h.OllamaService.Show(req)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to show: %v", err)
	}
	return resp, nil
}

func (h *WorkerHandler) Pull(ctx context.Context, req *workerpb.PullRequest) (*workerpb.PullResponse, error) {
	resp, err := h.OllamaService.Pull(req)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to pull: %v", err)
	}
	return resp, nil
}

func (h *WorkerHandler) Create(ctx context.Context, req *workerpb.CreateRequest) (*workerpb.CreateResponse, error) {
	resp, err := h.OllamaService.Create(req)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create: %v", err)
	}
	return resp, nil
}

func (h *WorkerHandler) Delete(ctx context.Context, req *workerpb.DeleteRequest) (*workerpb.DeleteResponse, error) {
	resp, err := h.OllamaService.Delete(req)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to delete: %v", err)
	}
	return resp, nil
}

func (h *WorkerHandler) HealthCheck(ctx context.Context, req *workerpb.HealthRequest) (*workerpb.HealthResponse, error) {
	return &workerpb.HealthResponse{Healthy: true, Details: "OK"}, nil
}
