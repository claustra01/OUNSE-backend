package api

import (
	"hackz-allo/db"
	"hackz-allo/utils"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func CreatePost(c echo.Context) error {

	type json struct {
		Title string `json:"title"`
		Body  string `json:"body"`
		User  string `json:"user_id"`
	}

	// クエリ展開
	o := new(json)
	if err := c.Bind(o); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// 投稿作成
	uuidObj, _ := uuid.NewUUID()
	post := new(db.Post)
	post.Id = uuidObj
	post.Title = o.Title
	post.Body = o.Body
	post.Time = utils.TimeToString(time.Now())
	post.UserId = o.User
	db.Psql.Create(&post)

	// 投稿時間を返す
	return c.String(http.StatusOK, utils.TimeToString(time.Now()))
}
