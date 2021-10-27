// бизнес логика

package service

import (
	"github.com/piklin/test_ozon_fintech/models"
	"github.com/piklin/test_ozon_fintech/pkg/repository"
)

//go:generate go run github.com/golang/mock/mockgen -source=service.go -destination=mocks/mock.go

type ShortURL interface {
	Create(url models.URLRequest) (string, error)
	GetFullURL(shortURL string) (string, error)
	GenerateShortURL() (string, error)
}

type Service struct {
	ShortURL
}

func NewService(r *repository.Repository) *Service {
	return &Service{
		ShortURL: NewShortURLService(r.ShortURL),
	}
}
