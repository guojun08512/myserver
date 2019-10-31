package server

import (
	"errors"
	"fmt"

	"myserver/pkg/config"
	"myserver/pkg/db"
	"myserver/pkg/logger"
	"myserver/web"

	"github.com/labstack/echo/v4"
)

var log = logger.WithNamespace("server")

func Start() {
	e := echo.New()
	// Routes
	web.SetupRoutes(e)
	dbClient := db.GetDbClient()
	if dbClient == nil {
		panic(errors.New("db client init failed"))
	}
	log.Fatal(e.Start(fmt.Sprintf(":%d", config.GetConfig().Port)))
	defer dbClient.Close()
}
