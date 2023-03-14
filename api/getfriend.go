package api

import (
	"hackz-allo/db"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetFriend(c echo.Context) error {

	type response struct {
		FriendList  []db.Friend
		RequestList []db.Friend
	}

	user := c.QueryParam("user_id")

	// フレンド情報取得
	obj := new(response)
	db.Psql.Where("user_id = ?", user).First(&obj.FriendList)
	db.Psql.Where("friend_id = ?", user).Where("is_request = ?", true).First(&obj.RequestList)

	return c.JSON(http.StatusOK, obj)
}
