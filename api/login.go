package api

import (
	"hackz-allo/database"
	"net/http"

	"github.com/labstack/echo/v4"
)

func LogIn(c echo.Context) error {

	type json struct {
		Id       string `json:"user_id"`
		Password string `json:"password"`
	}

	type response struct {
		Result  string
		Message string
	}

	db := database.Connect()
	obj := new(response)

	// クエリ展開
	o := new(json)
	if err := c.Bind(o); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	id := o.Id
	password := o.Password

	// ログイン判定
	array := []database.User{}
	db.Find(&array)
	for _, u := range array {
		if u.UserId == id {
			if u.Password == password {
				// 成功
				obj.Result = "OK"
				obj.Message = u.Id.String()
				return c.JSON(http.StatusOK, obj)
			} else {
				// パスワードが違う時
				obj.Result = "Failed"
				obj.Message = "パスワードが違います"
				return c.JSON(http.StatusOK, obj)
			}
		}
	}

	// ユーザーが見つからない時
	obj.Result = "Failed"
	obj.Message = "ユーザーが見つかりません"
	return c.JSON(http.StatusOK, obj)
}
