package auth

import (
	"github.com/labstack/echo/v4"
)

func AuthRouter(router *echo.Group) {
	router.POST("/login", login)
	router.POST("/user", register)
}
