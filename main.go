package main

import (
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"hackz-allo/api"
	"hackz-allo/db"
)

func main() {

	// 環境変数の読み込み
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	// インスタンス作成
	e := echo.New()

	// ミドルウェア設定
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// CORSの設定
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{os.Getenv("CLIENT_URL_LOCAL_1"), os.Getenv("CLIENT_URL_LOCAL_2"), os.Getenv("CLIENT_URL_REMOTE")},
		AllowMethods: []string{
			http.MethodGet,
			http.MethodPut,
			http.MethodPost,
			http.MethodDelete,
		},
		AllowCredentials: true,
	}))

	// DB接続
	db.Connect()

	// DB初期化 開発用
	e.GET("/api/initialization", api.Initialization)

	// 認証
	e.GET("/api/auth", api.Auth)
	e.POST("/api/login", api.LogIn)
	e.POST("/api/signup", api.SignUp)

	// ユーザー
	e.GET("api/getuser", api.GetUser)
	e.POST("api/sendfriend", api.SendFriend)
	e.DELETE("api/removefriend", api.RemoveFriend)
	e.GET("api/getfriend", api.GetFriend)

	// 投稿
	e.POST("api/createpost", api.CreatePost)
	e.PUT("api/editpost", api.EditPost)
	e.DELETE("api/deletepost", api.DeletePost)
	e.GET("api/gettimeline", api.GetTimeLine)

	// サーバー起動
	e.Logger.Fatal(e.Start(":8080"))
}
