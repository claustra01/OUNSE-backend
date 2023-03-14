package api

import (
	"hackz-allo/db"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Auth(c echo.Context) error {

	token := c.QueryParam("token")

	// 認証チェック
	array := []db.User{}
	db.Psql.Find(&array)
	for _, u := range array {
		if u.Id.String() == token {
			return c.String(http.StatusOK, "OK")
		}
	}
	return c.String(http.StatusOK, "Failed")
}
