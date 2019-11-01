package web

import (
	"fmt"
	"net/http"

	"myserver/web/middlewares"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func version(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}

func test(c echo.Context) error {
	fmt.Println("test")
	return c.NoContent(http.StatusOK)
}

// SetupRoutes 建立controller
func SetupRoutes(router *echo.Echo) {
	// Middleware
	router.Use(middleware.Logger())
	setupRecover(router)
	router.GET("/", version)
	{
		mwsNotBlocked := []echo.MiddlewareFunc{
			middlewares.AuthMiddleware(),
			middlewares.Accept(middlewares.AcceptOptions{
				DefaultContentTypeOffer: "application/json",
			}),
		}
		router.Group("/v1")
		router.GET("/test", test,  mwsNotBlocked...)
	}
}

func setupRecover(router *echo.Echo) {
	recoverMiddleware := middlewares.RecoverWithConfig(middlewares.RecoverConfig{
		StackSize: 10 << 10, // 10KB
	})
	router.Use(recoverMiddleware)
}
