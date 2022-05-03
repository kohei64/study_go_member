package router

import (
	
	"go-member/controller"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func Init() {


	e := echo.New()
	e.Use(middleware.CORS())
	e.Pre(controller.MethodOverride)

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// ルーター
	e.GET("/users",controller.GetUsers) //全ユーザー取得
	e.GET("/user/:id", controller.GetUser) //ユーザー取得
	e.POST("/user", controller.CreateUser) //ユーザー作成
	e.PUT("/user/:id", controller.UpdateUser) //ユーザー編集
	e.DELETE("/user/:id", controller.DeleteUser) //ユーザー削除

	e.Logger.Fatal(e.Start("localhost:8080"))
}

