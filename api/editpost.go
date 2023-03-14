package api

import (
	"hackz-allo/db"
	"hackz-allo/utils"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func EditPost(c echo.Context) error {

	type json struct {
		Id    string `json:"id"`
		Title string `json:"title"`
		Body  string `json:"body"`
	}

	// クエリ展開
	o := new(json)
	if err := c.Bind(o); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// 投稿更新
	p := new(db.Post)
	db.Psql.Where("id = ?", o.Id).First(&p)
	p.Title = o.Title
	p.Body = o.Body
	p.Time = utils.TimeToString(time.Now())
	db.Psql.Save(&p)

	// 更新時間を返す
	return c.String(http.StatusOK, utils.TimeToString(time.Now()))
}
