package server

import (
	"fmt"

	"myserver/pkg/config"
	"myserver/pkg/logger"
	"myserver/web"

	"github.com/labstack/echo/v4"
)

var log = logger.WithNamespace("server")

func Start() {
	e := echo.New()
	// Routes
	web.SetupRoutes(e)
	log.Fatal(e.Start(fmt.Sprintf(":%d", config.GetConfig().Port)))
}
