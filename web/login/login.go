package login

import (
	"errors"
	"github.com/labstack/echo"
	"myserver/pkg/crypto"
	"myserver/pkg/db"
	"myserver/pkg/logger"
	"net/http"
)

var log = logger.WithNamespace("web/login")

var (
	userNotFount = errors.New("User Not Found, Check Your Name, Password")
)

type loginInfo struct {
	Name string `json:"name"`
	PassWord string `json:"password"`
}



func login(c echo.Context) error {
	l := &loginInfo{}
	if err := c.Bind(l); err != nil {
		log.Errorf("login failed %v", err.Error())
		return c.String(http.StatusBadRequest, err.Error())
	}
	gorm := db.GetORM()
	u, err := db.GetUser(gorm, l.Name, l.PassWord)
	if err != nil {
		log.Errorf("login failed when GetUser: %v", err)
		return c.String(http.StatusBadRequest, userNotFount.Error())
	}
	token := crypto.CreateToken(u.ID, u.Roles)
	c.String(http.StatusOK, token)
}

func Router(router *echo.Group) {
	router.POST("/login", login)
}
