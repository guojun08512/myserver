package web

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func version(c echo.Context) error {
	fmt.Println(c.Request().Body)
	return c.NoContent(http.StatusOK)
}

func SetupRoutes(router *echo.Echo) {
	// Middleware
	// router.Use(middleware.Logger())
	// setupRecover(router)
	router.GET("/", version)
	// router.GET("/weixin", wxAuthorization)
	// router.GET("/code", wxresponse)
	// router.POST("/weixin", wxpost)
	// {
	// 	router.Group("v1")
	// }
}

// func setupRecover(router *echo.Echo) {
// 	recoverMiddleware := middlewares.RecoverWithConfig(middlewares.RecoverConfig{
// 		StackSize: 10 << 10, // 10KB
// 	})
// 	router.Use(recoverMiddleware)
// }
