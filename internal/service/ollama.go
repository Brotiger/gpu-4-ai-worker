package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"gpu-4-ai-worker/internal/config"
	"github.com/Brotiger/gpu-4-ai-worker/proto"
)

type OllamaService struct {
	Cfg *config.Config
}

func NewOllamaService(cfg *config.Config) *OllamaService {
	return &OllamaService{Cfg: cfg}
}

func (s *OllamaService) Post(endpoint string, req any, resp any) error {
	var bodyBytes []byte
	var err error
	if req != nil {
		bodyBytes, err = json.Marshal(req)
		if err != nil {
			return err
		}
	}
	ollamaURL := fmt.Sprintf("http://%s:%s/api/%s", s.Cfg.OllamaDomain, s.Cfg.OllamaPort, endpoint)
	httpResp, err := http.Post(ollamaURL, "application/json", bytes.NewReader(bodyBytes))
	if err != nil {
		return err
	}
	defer httpResp.Body.Close()
	return json.NewDecoder(httpResp.Body).Decode(resp)
}

func (s *OllamaService) Generate(req *proto.GenerateRequest) (*proto.GenerateResponse, error) {
	ollamaReq := map[string]interface{}{
		"model":  req.Model,
		"prompt": req.Prompt,
		"stream": req.Stream,
	}
	var ollamaResp struct {
		Response string `json:"response"`
		Done     bool   `json:"done"`
	}
	err := s.Post("generate", ollamaReq, &ollamaResp)
	if err != nil {
		return nil, err
	}
	return &proto.GenerateResponse{
		Response: ollamaResp.Response,
		Done:     ollamaResp.Done,
	}, nil
}

func (s *OllamaService) Tags() (*proto.TagsResponse, error) {
	var ollamaResp struct {
		Models []string `json:"models"`
	}
	err := s.Post("tags", nil, &ollamaResp)
	if err != nil {
		return nil, err
	}
	return &proto.TagsResponse{Models: ollamaResp.Models}, nil
}

func (s *OllamaService) Show(req *proto.ShowRequest) (*proto.ShowResponse, error) {
	ollamaReq := map[string]interface{}{"model": req.Model}
	var ollamaResp struct {
		Model   string            `json:"model"`
		Details map[string]string `json:"details"`
	}
	err := s.Post("show", ollamaReq, &ollamaResp)
	if err != nil {
		return nil, err
	}
	return &proto.ShowResponse{Model: ollamaResp.Model, Details: ollamaResp.Details}, nil
}

func (s *OllamaService) Pull(req *proto.PullRequest) (*proto.PullResponse, error) {
	ollamaReq := map[string]interface{}{"name": req.Name}
	var ollamaResp struct {
		Status string `json:"status"`
	}
	err := s.Post("pull", ollamaReq, &ollamaResp)
	if err != nil {
		return nil, err
	}
	return &proto.PullResponse{Status: ollamaResp.Status}, nil
}

func (s *OllamaService) Create(req *proto.CreateRequest) (*proto.CreateResponse, error) {
	ollamaReq := map[string]interface{}{"name": req.Name, "modelfile": req.Modelfile}
	var ollamaResp struct {
		Status string `json:"status"`
	}
	err := s.Post("create", ollamaReq, &ollamaResp)
	if err != nil {
		return nil, err
	}
	return &proto.CreateResponse{Status: ollamaResp.Status}, nil
}

func (s *OllamaService) Delete(req *proto.DeleteRequest) (*proto.DeleteResponse, error) {
	ollamaReq := map[string]interface{}{"model": req.Model}
	var ollamaResp struct {
		Status string `json:"status"`
	}
	err := s.Post("delete", ollamaReq, &ollamaResp)
	if err != nil {
		return nil, err
	}
	return &proto.DeleteResponse{Status: ollamaResp.Status}, nil
} 