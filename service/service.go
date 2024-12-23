package service

import (
	"high-traffic-practice/config"
	auth "high-traffic-practice/gRPC/proto"
	"high-traffic-practice/repository"
)

type Service struct {
	cfg *config.Config

	repository *repository.Repository
}

func NewService(cfg *config.Config, repository *repository.Repository) (*Service, error) {
	r := &Service{cfg: cfg, repository: repository}

	return r, nil
}

func (s *Service) CreateAuth(username string) (*auth.AuthData, error) {
	return s.repository.CreateAuth(username)
}
