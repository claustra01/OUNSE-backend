package api

import (
	"hackz-allo/database"
	"net/http"

	"github.com/labstack/echo/v4"
)

func DeletePost(c echo.Context) error {

	db := database.Connect()
	id := c.QueryParam("id")

	// 投稿削除
	p := new(database.Post)
	db.Where("id = ?", id).First(&p)
	db.Where("id = ?", id).Delete(&database.Post{})

	// 削除した投稿を返す
	database.Close(db)
	return c.JSON(http.StatusOK, p)
}
