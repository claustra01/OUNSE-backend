package api

import (
	"hackz-allo/db"
	"hackz-allo/utils"
	"net/http"

	"github.com/labstack/echo/v4"
	"golang.org/x/exp/slices"
)

func SendFriend(c echo.Context) error {

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

	// レコード取得
	recu := new(db.Friend)
	db.Psql.Where("user_id = ?", user).First(&recu)
	recf := new(db.Friend)
	db.Psql.Where("user_id = ?", user).First(&recf)

	// 追加して保存
	if slices.Contains(recf.RequestUser, user) {
		recu.FriendUser = append(recu.FriendUser, friend)
		recf.FriendUser = append(recf.FriendUser, user)
		recf.RequestUser = utils.RemoveFromSlice(recf.RequestUser, user)
	} else {
		recu.RequestUser = append(recu.RequestUser, friend)
	}
	db.Psql.Save(&recu)
	db.Psql.Save(&recf)

	return c.JSON(http.StatusOK, nil)
}
