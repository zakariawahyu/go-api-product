package main

import (
	"github.com/zakariawahyu/go-api-product/config"
	"github.com/zakariawahyu/go-api-product/internal/server"
	"github.com/zakariawahyu/go-api-product/pkg/logger"
	"github.com/zakariawahyu/go-api-product/pkg/mysql"
	"log"
)

// @title Product Api Services
// @version 1.0.0
// @description Product REST API
// @termsOfService http://swagger.io/terms/

// @contact.name Zakaria Wahyu
// @contact.url https://github.com/zakariawahyu
// @contact.email zakarianur6@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:7070
// @BasePath /api/v1/product

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
