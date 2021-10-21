package main

import (
	"flag"
	log "github.com/sirupsen/logrus"
	_ "github.com/lib/pq"
	forum "github.com/piklin/test_ozon_fintech/pkg"
	"github.com/piklin/test_ozon_fintech/pkg/handler"
	"github.com/piklin/test_ozon_fintech/pkg/repository"
	"github.com/piklin/test_ozon_fintech/pkg/service"
)

func main() {
	var databaseType string				//redis or postgres
	flag.StringVar(&databaseType, "db", "postgres", "database type")
	flag.Parse()

	db := repository.NewDatabase(databaseType)

	repository := repository.NewRepository(db)
	services := service.NewService(repository)
	handlers := handler.NewHandler(services)

	server := new(forum.Server)
	if error := server.Run(handlers.InitRoutes()); error != nil {
		log.WithFields(log.Fields{
			"package": 		"main",
			"function":		"main",
			"error":			error,
		}).Fatal("Http server starting error")
	}
}
