// работа с базой данных

package repository

import (
	//"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
)

//go:generate go run github.com/golang/mock/mockgen -source=repository.go -destination=mocks/mock.go

type ShortURL interface {
	Create(generatedURL, full_url string) error
	SearchShortURL(shortURL string) (string, error)
	SearchFullURL(fullURL string) (string, error)
}

type Repository struct {
	ShortURL
}

func NewRepository(db Database) *Repository {
	if db.DBType == Postgres {
		return &Repository{
			ShortURL: NewShortURLPostgres(db.Postgres),
		}
	} else if db.DBType == Redis {
		return &Repository{
			ShortURL: NewShortURLRedis(db.Redis),
		}
	} else {
		log.WithFields(log.Fields{
			"package":  "repository",
			"function": "NewRepository",
		}).Fatal("Unknown database")
		return &Repository{}
	}
}
