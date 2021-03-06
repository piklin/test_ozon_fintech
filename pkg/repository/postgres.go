package repository

import (
  "github.com/jmoiron/sqlx"
  "fmt"
	log "github.com/sirupsen/logrus"
)

type PostgresConfig struct {
  Host        string
  Port        string
  Username    string
  Password    string
  DBName      string
  SSLMode     string
}

func NewPostgresDB(config PostgresConfig) (*sqlx.DB, error) {
  db, error := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
    config.Host, config.Port, config.Username, config.DBName, config.Password, config.SSLMode))
  if error != nil {
    log.WithFields(log.Fields{
			"package": 		"repository",
			"function":		"NewPostgresDB",
			"error":			error,
		}).Fatal("Postgres database open error. ")
    return nil, error
  }

  error = db.Ping()
  if error != nil {
    log.WithFields(log.Fields{
			"package": 		"repository",
			"function":		"NewPostgresDB",
			"error":			error,
		}).Fatal("Postgres ping error. ")
    return nil, error
  }

  return db, nil
}
