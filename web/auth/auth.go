package auth

import (
	"errors"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"myserver/pkg/crypto"
	"myserver/pkg/db"
	"myserver/pkg/logger"
	"myserver/web/jsonapi"
	"net/http"
)

var log = logger.WithNamespace("web/login")

var (
	userNotFount = errors.New("User Not Found, Check Your Name, Password")
)

func login(c echo.Context) error {
	l := &userInfo{}
	if err := c.Bind(l); err != nil {
		log.Errorf("login failed %v", err.Error())
		return jsonapi.ResponseWithJson(c, http.StatusBadRequest, err)
	}
	gorm := db.GetORM()
	u, err := db.GetUser(gorm, l.Name, l.PassWord)
	if err != nil {
		log.Errorf("login failed when GetUser: %v", err)
		return jsonapi.ResponseWithJson(c, http.StatusBadRequest, err)
	}
	var roles []string
	err = json.Unmarshal([]byte(u.Roles), &roles)
	if err != nil {
		log.Errorf("login failed when json parsed: %v", err)
		return jsonapi.ResponseWithJson(c, http.StatusInternalServerError, err)
	}
	token ,err := crypto.CreateToken(u.ID, roles)
	if err != nil {
		return jsonapi.ResponseWithJson(c, http.StatusInternalServerError, err)
	}
	return jsonapi.ResponseWithJson(c, http.StatusOK,  map[string]interface{}{"token": token })
}

func Router(router *echo.Group) {
	router.POST("/login", login)
}
