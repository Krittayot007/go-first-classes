package service

import (
	"context"

	"github.com/SalhSThai/go-first-classes/internal/client"
	"github.com/SalhSThai/go-first-classes/internal/repository"
)

type serviceImpl struct {
	repo   repository.Repository
	client client.Client
}

type Service interface {
	ServiceRegisterPost(ctx context.Context) error
}

func New(r repository.Repository, c client.Client) (Service, error) {
	return &serviceImpl{
		repo:   r,
		client: c,
	}, nil
}
