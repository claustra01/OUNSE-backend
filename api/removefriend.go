package api

import (
	"hackz-allo/database"
	"hackz-allo/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

func RemoveFriend(c echo.Context) error {

	type json struct {
		User   string `json:"user_id"`
		Friend string `json:"friend_id"`
	}

	// クエリ展開
	o := new(json)
	if err := c.Bind(o); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	user := o.User
	friend := o.Friend

	db := database.Connect()

	// レコード取得
	rec := new(database.Friend)
	db.Where("user_id = ?", user).First(&rec)
	r := rec.RequestUser
	f := rec.FriendUser

	// 削除して保存
	rec.RequestUser = utils.RemoveFromSlice(r, friend)
	rec.FriendUser = utils.RemoveFromSlice(f, friend)
	db.Save(&rec)

	database.Close(db)
	return c.JSON(http.StatusOK, nil)
}
