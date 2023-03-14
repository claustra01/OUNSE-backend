package api

import (
	"hackz-allo/db"
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

	// 削除
	rec := new(db.Friend)
	db.Psql.Where("user_id = ?", user).Where("friend_id = ?", friend).First(&rec)
	db.Psql.Delete(&rec)

	return c.JSON(http.StatusOK, nil)
}
