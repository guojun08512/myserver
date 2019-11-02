package auth

import (
	"github.com/labstack/echo/v4"
	"myserver/pkg/db"
	"net/http"
)

type userInfo struct {
	ID string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	PassWord string `json:"password,omitempty"`
	Email string `json:"email,omitempty"`
	Sex string `json:"sex,omitempty"`
	Phone string `json:"phone,omitempty"`
	Role string `json:"role"`
}

func register(c echo.Context) error {
	var user userInfo
	err := c.Bind(&user)
	if err != nil {
		log.Errorf("register faild when bind: %v", err)
		return c.String(http.StatusBadRequest, err.Error())
	}
	u := db.CreateUser(db.GetORM(), user.Name, user.PassWord, user.Email, user.Role, user.Phone)
	log.Infof("register success userID(%s)", u.ID)
	return c.String(http.StatusOK, u.ID)
}

