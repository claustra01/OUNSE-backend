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

	// ユーザー存在確認
	arr := []db.User{}
	db.Psql.Find(&arr)
	exist := false
	for _, u := range arr {
		if u.UserId == friend {
			exist = true
		}
	}
	if !exist {
		return c.String(http.StatusOK, "ユーザーが見つかりません")
	}

	// 重複リクエスト確認
	arr2 := []db.Friend{}
	db.Psql.Find(&arr2)
	req := false
	for _, f := range arr2 {
		if f.UserId == user && f.FriendId == friend {
			req = true
		}
	}
	if req {
		return c.String(http.StatusOK, "既にフレンドかリクエスト済みです")
	}

	// 既にリクエストを受け取っているか
	req = false
	db.Psql.Where("user_id = ?", friend).Find(&arr2)
	for _, f := range arr2 {
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

	return c.String(http.StatusOK, "OK")
}
