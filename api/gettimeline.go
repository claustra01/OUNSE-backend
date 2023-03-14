package api

import (
	"hackz-allo/database"
	"hackz-allo/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetTimeLine(c echo.Context) error {

	type response struct {
		posts []database.Post
	}

	db := database.Connect()
	user := c.QueryParam("user_id")

	obj := new(response)
	obj.posts = []database.Post{}

	// フレンド情報取得
	rec := new(database.Friend)
	db.Where("user_id = ?", user).First(&rec)
	friends := rec.FriendUser

	// 投稿取得
	p := []database.Post{}
	db.Where("user_id = ?", user).Find(&p)
	for _, f := range friends {
		q := []database.Post{}
		db.Where("user_id = ?", f).Find(&q)
		p = append(p, q...)
	}

	// 投稿ソート
	obj.posts = utils.SortPost(p, 48)

	return c.JSON(http.StatusOK, obj)
}
