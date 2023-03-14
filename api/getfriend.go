package api

import (
	"hackz-allo/database"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetFriend(c echo.Context) error {

	type response struct {
		Friend  []string
		Request []string
	}

	db := database.Connect()
	user := c.QueryParam("user_id")

	return c.JSON(http.StatusOK, nil)
}
