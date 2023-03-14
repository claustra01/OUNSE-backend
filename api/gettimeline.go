package api

import (
	"hackz-allo/db"
	"hackz-allo/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetTimeLine(c echo.Context) error {

	user := c.QueryParam("user_id")

	// フレンド情報取得
	rec := []db.Friend{}
	db.Psql.Where("user_id = ?", user).Where("is_request = ?", false).Find(&rec)

	// 投稿取得
	p := []db.Post{}
	db.Psql.Where("user_id = ?", user).Find(&p)
	for _, r := range rec {
		q := []db.Post{}
		db.Psql.Where("user_id = ?", r.FriendId).Find(&q)
		p = append(p, q...)
	}

	// 投稿ソート
	p = utils.SortPost(p, 48)

	return c.JSON(http.StatusOK, p)
}
