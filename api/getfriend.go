package api

import (
	"hackz-allo/database"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetFriend(c echo.Context) error {

	db := database.Connect()
	user := c.QueryParam("user_id")

	// フレンド情報取得
	obj := new(database.Friend)
	db.Where("user_id = ?", user).First(&obj)

	return c.JSON(http.StatusOK, obj)
}
