package api

import (
	"hackz-allo/db"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetUser(c echo.Context) error {

	type response struct {
		UserId string
		Name   string
	}

	token := c.QueryParam("token")

	// ユーザー情報取得
	array := []db.User{}
	db.Psql.Find(&array)
	for _, u := range array {
		if u.Id.String() == token {
			obj := new(response)
			obj.UserId = u.UserId
			obj.Name = u.Name
			return c.JSON(http.StatusOK, obj)
		}
	}
	return c.JSON(http.StatusOK, nil)
}
