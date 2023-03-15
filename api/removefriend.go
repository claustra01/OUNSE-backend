package api

import (
	"hackz-allo/db"
	"net/http"

	"github.com/labstack/echo/v4"
)

func RemoveFriend(c echo.Context) error {

	// クエリ展開
	user := c.QueryParam("user_id")
	friend := c.QueryParam("friend_id")

	// 削除
	rec := new(db.Friend)
	db.Psql.Where("user_id = ?", user).Where("friend_id = ?", friend).First(&rec)
	db.Psql.Where("user_id = ?", user).Where("friend_id = ?", friend).Delete(&rec)
	db.Psql.Where("user_id = ?", friend).Where("friend_id = ?", user).First(&rec)
	db.Psql.Where("user_id = ?", friend).Where("friend_id = ?", user).Delete(&rec)

	return c.JSON(http.StatusOK, nil)
}
