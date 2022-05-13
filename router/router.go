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
	e.POST("/login", controller.Login)

	r:=e.Group("/restricted")
	r.Use(middleware.JWTWithConfig(controller.Config))
	
	r.GET("/users",controller.GetUsers) //全ユーザー取得
	r.GET("/user/:id", controller.GetUser) //ユーザー取得
	r.POST("/user", controller.CreateUser) //ユーザー作成
	r.PUT("/user/:id", controller.UpdateUser) //ユーザー編集
	r.DELETE("/user/:id", controller.DeleteUser) //ユーザー削除

	e.Logger.Fatal(e.Start("localhost:8080"))
}

