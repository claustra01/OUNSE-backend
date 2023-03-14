package api

import (
	"hackz-allo/db"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Initialization(c echo.Context) error {

	db.Psql.Exec("DROP TABLE IF EXISTS users")
	db.Psql.Exec("DROP TABLE IF EXISTS posts")
	db.Psql.Exec("DROP TABLE IF EXISTS friends")

	db.Psql.AutoMigrate(db.User{})
	db.Psql.AutoMigrate(db.Post{})
	db.Psql.AutoMigrate(db.Friend{})

	return c.String(http.StatusOK, "Initializaton")
}
