package repository

import (
  log "github.com/sirupsen/logrus"
  "github.com/go-redis/redis"
)

type RedisConfig struct {
  Host        string
  Port        string
  Password    string
  DB          int
}

func NewRedisDB(config RedisConfig) (*redis.Client, error) {
  client := redis.NewClient(&redis.Options{
		Addr:     config.Host + ":" + config.Port,
		Password: config.Password,
		DB:       config.DB,
	})

  _, error := client.Ping().Result()

  if error != nil {
    log.WithFields(log.Fields{
			"package": 		"repository",
			"function":		"NewRedisDB",
			"error":			error,
		}).Fatal("Redis database open error. ")
    return nil, error
  }
  return client, nil
}
