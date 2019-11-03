package jsonapi

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type Response struct {
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func wrapResponse(code int, data interface{}) interface{} {
	switch code {
	case http.StatusOK:
		return Response{
			Msg: "success!!",
			Data: data,
		}
	case http.StatusInternalServerError:
		return Response{
			Msg: "Internal server error!!",
			Data: data,
		}
	case http.StatusBadRequest:
		return Response{
			Msg: "Bad request!!",
			Data: data,
		}
	case http.StatusConflict:
		return Response{
			Msg: "Conflict!!",
			Data: data,
		}
	case http.StatusMethodNotAllowed:
		return Response{
			Msg: "Method Not Allowed!!",
			Data: data,
		}
	case http.StatusNotFound:
		return Response{
			Msg: "Not Fount!!",
			Data: data,
		}
	}
	return nil
}

func ResponseWithJson(c echo.Context, code int, payload interface{}) error {
	return c.JSON(code, wrapResponse(code, payload))
}
