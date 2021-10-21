package repository

import (
  "github.com/go-redis/redis"
  log "github.com/sirupsen/logrus"
)

type ShortURLRedis struct {
	db *redis.Client
}

func NewShortURLRedis(db *redis.Client) *ShortURLRedis {
	return &ShortURLRedis{db: db}
}

func (r *ShortURLRedis) SearchShortURL(shortURL string) (string, error) {
  fullURL, error := r.db.Get(shortURL).Result()
  if error == redis.Nil {
    return "", nil
  } else if error != nil {
    log.WithFields(log.Fields{
      "package": 		"repository",
      "struct":			"ShortURLRedis",
      "function":		"SearchShortURL",
      "error":			error,
    }).Error("Get error")
    return "", error
  }
	return fullURL, nil
}

func (r *ShortURLRedis) SearchFullURL(fullURL string) (string, error) {
  shortURL, error := r.db.Get(fullURL).Result()
  if error == redis.Nil {
    return "", nil
  } else if error != nil {
    log.WithFields(log.Fields{
      "package": 		"repository",
      "struct":			"ShortURLRedis",
      "function":		"SearchFullURL",
      "error":			error,
    }).Error("Get error")
    return "", error
  }
  return shortURL, nil
}

func (r *ShortURLRedis) Create(generatedURL, fullURL string) error {
  if error := r.db.Set(generatedURL, fullURL, 0).Err(); error != nil {
    log.WithFields(log.Fields{
      "package": 		"repository",
      "struct":			"ShortURLRedis",
      "function":		"Create",
      "error":			error,
    }).Error("Set error")
    return error
  }
  if error := r.db.Set(fullURL, generatedURL, 0).Err(); error != nil {
    log.WithFields(log.Fields{
      "package": 		"repository",
      "struct":			"ShortURLRedis",
      "function":		"Create",
      "error":			error,
    }).Error("Set error")
    return error
  }
	return nil
}
