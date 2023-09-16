package main

import (
	"github.com/zakariawahyu/go-api-product/config"
	"github.com/zakariawahyu/go-api-product/internal/server"
	"github.com/zakariawahyu/go-api-product/pkg/logger"
	"github.com/zakariawahyu/go-api-product/pkg/mysql"
	"log"
)

func main() {
	cfg, err := config.InitConfig()
	if err != nil {
		log.Fatal(err)
	}

	appLogger := logger.NewAppLogger(cfg)
	appLogger.InitLogger()

	db, err := mysql.NewDBConnection(cfg)
	if err != nil {
		appLogger.Fatal(err)
	}

	s := server.NewServer(appLogger, cfg, db)
	appLogger.Fatal(s.Run())
}
