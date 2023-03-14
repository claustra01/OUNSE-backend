package api

import (
	"hackz-allo/db"
	"net/http"

	"github.com/labstack/echo/v4"
)

func DeletePost(c echo.Context) error {

	id := c.QueryParam("id")

	// 投稿削除
	p := new(db.Post)
	db.Psql.Where("id = ?", id).First(&p)
	db.Psql.Where("id = ?", id).Delete(&db.Post{})

	// 削除した投稿を返す
	return c.JSON(http.StatusOK, p)
}
