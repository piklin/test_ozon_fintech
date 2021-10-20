package service

import (
	"math/rand"
	"time"

	"github.com/piklin/test_ozon_fintech/models"
	"github.com/piklin/test_ozon_fintech/pkg/repository"
)

const (
	symbols     = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_"
	shortURTLen = 10
	// HARDCODE!!!!! К сожалению, красивое и элегантное решение для хоста не смог придумать. Но, думаю, с опытом это придет
	host 				= "localhost:8080"
)

type ShortURLService struct {
	repository repository.ShortURL
}

func NewShortURLService(repository repository.ShortURL) *ShortURLService {
	return &ShortURLService{repository: repository}
}


func (s *ShortURLService) GetFullURL(shortURL string) (string, error) {
	fullURL, error := s.repository.SearchShortURL(shortURL)
	if error != nil {
		return "", error
	}
	return fullURL, nil
}


func (s *ShortURLService) Create(url models.URL) (string, error) {
	shortURL, error := s.repository.SearchFullURL(url.URL)
	if error != nil {
		return "", error
	}

	if shortURL != "" {
		return host + "/" + shortURL, nil
	}

	generatedURL, error := s.generateShortURL()
	if error != nil {
		return "", error
	}

	return host + "/" + generatedURL, s.repository.Create(generatedURL, url.URL)
}


func (s *ShortURLService) generateShortURL() (string, error) {
	shortURL := make([]byte, shortURTLen)

	rand.Seed(time.Now().UnixNano())

	for {
		for i := range shortURL {
			shortURL[i] = symbols[rand.Intn(len(symbols))]
		}

		result, error := s.repository.SearchShortURL(string(shortURL))
		if error != nil {
			return "", error
		}
		if result == "" {
			break
		}
	}

	return string(shortURL), nil
}
