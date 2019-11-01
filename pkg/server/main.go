package server

import (
	"fmt"

	"myserver/pkg/config"
	"myserver/pkg/db"
	"myserver/pkg/logger"
	"myserver/web"

	"github.com/labstack/echo/v4"
)

var log = logger.WithNamespace("server")

// Start 启动服务
func Start() {
	e := echo.New()
	// Routes
	web.SetupRoutes(e)
	dbName := config.GetConfig().DB.DBName
	db.NewDB(dbName)
	e.Start(fmt.Sprintf(":%d", config.GetConfig().Port))
	defer db.CloseORM()
}
