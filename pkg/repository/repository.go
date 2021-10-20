// работа с базой данных

package repository

import (
	"github.com/jmoiron/sqlx"
)

type ShortURL interface {
	Create(generatedURL, full_url string) error
	SearchShortURL(shortURL string) (string, error)
	SearchFullURL(fullURL string) (string, error)
}

type Repository struct {
	ShortURL
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		ShortURL: NewShortURLPostgres(db),
	}
}
