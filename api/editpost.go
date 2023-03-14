package api

import (
	"hackz-allo/database"
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

	db := database.Connect()

	// 投稿更新
	p := new(database.Post)
	db.Where("id = ?", o.Id).First(&p)
	p.Title = o.Title
	p.Body = o.Body
	p.Time = utils.TimeToString(time.Now())
	db.Save(&p)

	// 更新時間を返す
	database.Close(db)
	return c.String(http.StatusOK, utils.TimeToString(time.Now()))
}
