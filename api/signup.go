package api

import (
	"hackz-allo/db"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func SignUp(c echo.Context) error {

	type json struct {
		Id       string `json:"user_id"`
		Name     string `json:"name"`
		Password string `json:"password"`
	}

	type response struct {
		Result  string
		Message string
	}

	// クエリ展開
	o := new(json)
	if err := c.Bind(o); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	id := o.Id
	name := o.Name
	password := o.Password

	obj := new(response)

	// ユーザー名重複チェック
	array := []db.User{}
	db.Psql.Find(&array)
	dup := false
	for _, u := range array {
		if u.UserId == id {
			dup = true
		}
	}
	if dup {
		obj.Result = "Failed"
		obj.Message = "ID(" + id + ")は既に使われています"
		return c.JSON(http.StatusOK, obj)
	}

	// ユーザー登録
	uuidObj, _ := uuid.NewUUID()
	user := new(db.User)
	user.Id = uuidObj
	user.UserId = id
	user.Name = name
	user.Password = password
	db.Psql.Create(&user)

	obj.Result = "OK"
	obj.Message = "ID(" + id + ") is registered!"
	return c.JSON(http.StatusOK, obj)
}
