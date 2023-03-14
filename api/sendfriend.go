package api

import (
	"hackz-allo/db"
	"net/http"

	"github.com/labstack/echo/v4"
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

	// 既にリクエストを受け取っているか
	req := false
	arr := []db.Friend{}
	db.Psql.Where("user_id = ?", friend).Find(&arr)
	for _, f := range arr {
		if f.FriendId == user {
			req = true
		}
	}

	// フレンド追加
	u := new(db.Friend)
	u.UserId = user
	u.FriendId = friend
	if req {
		u.IsRequest = false
		db.Psql.Create(&u)
		f := new(db.Friend)
		db.Psql.Where("user_id = ?", friend).Where("friend_id = ?", user).First(&f)
		f.IsRequest = false
		db.Psql.Where("user_id = ?", friend).Where("friend_id = ?", user).Save(&f)
	} else {
		u.IsRequest = true
		db.Psql.Create(&u)
	}

	return c.JSON(http.StatusOK, nil)
}
