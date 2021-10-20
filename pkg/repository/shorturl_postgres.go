package repository

import (
	"database/sql"
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/jmoiron/sqlx"
)

type ShortURLPostgres struct {
	db *sqlx.DB
}

func NewShortURLPostgres(db *sqlx.DB) *ShortURLPostgres {
	return &ShortURLPostgres{db: db}
}

func (r *ShortURLPostgres) SearchShortURL(shortURL string) (string, error) {
	var fullURL string
	query := fmt.Sprint("SELECT full_url FROM short_urls WHERE short_url = $1")
	row := r.db.QueryRow(query, shortURL)
	error := row.Scan(&fullURL)

	if error == sql.ErrNoRows {
		return "", nil
	} else if error != nil {
		log.WithFields(log.Fields{
			"package": 		"repository",
			"struct":			"ShortURLPostgres",
			"function":		"SearchShortURL",
			"error":			error,
		}).Error("Sqxl Scan error")
		return "", error
	}
	return fullURL, nil
}

func (r *ShortURLPostgres) SearchFullURL(fullURL string) (string, error) {
	var shortURL string
	query := fmt.Sprint("SELECT short_url FROM short_urls WHERE full_url = $1")
	row := r.db.QueryRow(query, fullURL)
	error := row.Scan(&shortURL)

	if error == sql.ErrNoRows {
		return "", nil
	} else if error != nil {
		log.WithFields(log.Fields{
			"package": 		"repository",
			"struct":			"ShortURLPostgres",
			"function":		"SearchFullURL",
			"error":			error,
		}).Error("Sqxl Scan error")
		return "", error
	}
	return shortURL, nil
}

func (r *ShortURLPostgres) Create(generatedURL, fullURL string) error {
	var id int
	query := fmt.Sprint("INSERT INTO short_urls (short_url, full_url) VALUES ($1, $2) RETURNING ID")
	row := r.db.QueryRow(query, generatedURL, fullURL)
	error := row.Scan(&id)
	if error != nil {
		log.WithFields(log.Fields{
			"package": 		"repository",
			"struct":			"ShortURLPostgres",
			"function":		"Create",
			"error":			error,
		}).Error("Sqxl Scan error")
	}
	return error
}
