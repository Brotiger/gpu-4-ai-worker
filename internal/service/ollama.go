package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gpu-4-ai-worker/internal/config"
	"net/http"

	workerpb "github.com/Brotiger/gpu-4-ai-worker/proto"
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

func (s *OllamaService) Generate(req *workerpb.GenerateRequest) (*workerpb.GenerateResponse, error) {
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
	return &workerpb.GenerateResponse{
		Response: ollamaResp.Response,
		Done:     ollamaResp.Done,
	}, nil
}

func (s *OllamaService) Tags() (*workerpb.TagsResponse, error) {
	var ollamaResp struct {
		Models []string `json:"models"`
	}
	err := s.Post("tags", nil, &ollamaResp)
	if err != nil {
		return nil, err
	}
	return &workerpb.TagsResponse{Models: ollamaResp.Models}, nil
}

func (s *OllamaService) Show(req *workerpb.ShowRequest) (*workerpb.ShowResponse, error) {
	ollamaReq := map[string]interface{}{"model": req.Model}
	var ollamaResp struct {
		Model   string            `json:"model"`
		Details map[string]string `json:"details"`
	}
	err := s.Post("show", ollamaReq, &ollamaResp)
	if err != nil {
		return nil, err
	}
	return &workerpb.ShowResponse{Model: ollamaResp.Model, Details: ollamaResp.Details}, nil
}

func (s *OllamaService) Pull(req *workerpb.PullRequest) (*workerpb.PullResponse, error) {
	ollamaReq := map[string]interface{}{"name": req.Name}
	var ollamaResp struct {
		Status string `json:"status"`
	}
	err := s.Post("pull", ollamaReq, &ollamaResp)
	if err != nil {
		return nil, err
	}
	return &workerpb.PullResponse{Status: ollamaResp.Status}, nil
}

func (s *OllamaService) Create(req *workerpb.CreateRequest) (*workerpb.CreateResponse, error) {
	ollamaReq := map[string]interface{}{"name": req.Name, "modelfile": req.Modelfile}
	var ollamaResp struct {
		Status string `json:"status"`
	}
	err := s.Post("create", ollamaReq, &ollamaResp)
	if err != nil {
		return nil, err
	}
	return &workerpb.CreateResponse{Status: ollamaResp.Status}, nil
}

func (s *OllamaService) Delete(req *workerpb.DeleteRequest) (*workerpb.DeleteResponse, error) {
	ollamaReq := map[string]interface{}{"model": req.Model}
	var ollamaResp struct {
		Status string `json:"status"`
	}
	err := s.Post("delete", ollamaReq, &ollamaResp)
	if err != nil {
		return nil, err
	}
	return &workerpb.DeleteResponse{Status: ollamaResp.Status}, nil
}
