package repository

import (
	"github.com/go-redis/redis"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
)

const (
	Unknown = iota
	Postgres
	Redis
)

type Database struct {
	Postgres *sqlx.DB
	Redis    *redis.Client
	DBType   int
}

func NewDatabase(dbType string) Database {
	if dbType == "postgres" {
		db, error := NewPostgresDB(PostgresConfig{
			Host:     "localhost",
			Port:     "5432",
			Username: "docker",
			Password: "docker",
			DBName:   "docker",
			SSLMode:  "disable",
		})
		if error != nil {
			log.WithFields(log.Fields{
				"package":  "repository",
				"function": "NewDatabase",
				"error":    error,
			}).Fatal("Postgres database initialize error")
		}
		return Database{
			Postgres: db,
			DBType:   Postgres,
		}
	} else if dbType == "redis" {
		db, error := NewRedisDB(RedisConfig{
			Host:     "localhost",
			Port:     "6379",
			Password: "",
			DB:       0,
		})
		if error != nil {
			log.WithFields(log.Fields{
				"package":  "repository",
				"function": "NewDatabase",
				"error":    error,
			}).Fatal("Redis database initialize error")
		}
		return Database{
			Redis:  db,
			DBType: Redis,
		}
	} else {
		log.WithFields(log.Fields{
			"package":  "repository",
			"function": "NewDatabase",
		}).Fatal("Unknown database")
		return Database{
			DBType: Unknown,
		}
	}
}
