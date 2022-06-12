package router

import (
	"net/http"
	"os"
	"study_go_member/controller"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func Init() {


	e := echo.New()
	// e.Use(middleware.CORS())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
	}))

	e.Pre(controller.MethodOverride)

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// ルーター
	e.POST("/login", controller.Login)

	r:=e.Group("/restricted")
	r.Use(middleware.JWTWithConfig(controller.Config))

	e.POST("/user", controller.CreateUser) //サインアップ&ユーザー作成
	e.POST("/login",controller.Login)//ログイン
	e.GET("/users",controller.GetUsers) //全ユーザー取得
	e.GET("/user/:id", controller.GetUser) //ユーザー取得
	e.PUT("/user/:id", controller.UpdateUser) //ユーザー編集
	e.DELETE("/user/:id", controller.DeleteUser) //ユーザー削除

	e.GET("/health",func(c echo.Context) error {
		return c.String(http.StatusOK, "health ok")
	})

	e.Logger.Fatal(e.Start(":"+os.Getenv("PORT")))
	// e.Logger.Fatal(e.Start(":8080"))
}

