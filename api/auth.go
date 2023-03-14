package api

import (
	"hackz-allo/database"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Auth(c echo.Context) error {

	db := database.Connect()
	token := c.QueryParam("token")

	// 認証チェック
	array := []database.User{}
	db.Find(&array)
	for _, u := range array {
		if u.Id.String() == token {
			database.Close(db)
			return c.String(http.StatusOK, "OK")
		}
	}
	database.Close(db)
	return c.String(http.StatusOK, "Failed")
}
