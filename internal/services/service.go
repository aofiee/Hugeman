package services

import "hugeman/internal/core/ports"

// Service struct
type Service struct {
	repo ports.Repository
}

// New func
func New(repo ports.Repository) *Service {
	return &Service{
		repo: repo,
	}
}
