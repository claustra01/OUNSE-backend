package api

import (
	"hackz-allo/db"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetFriend(c echo.Context) error {

	user := c.QueryParam("user_id")

	// フレンド情報取得
	obj := new(db.Friend)
	db.Psql.Where("user_id = ?", user).First(&obj)

	return c.JSON(http.StatusOK, obj)
}
