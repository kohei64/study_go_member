package controller

import (
	"net/http"
	"strconv"

	"go-member/model"

	"github.com/labstack/echo"
	"golang.org/x/crypto/bcrypt"
)

// JSONからRenderに変える

// メソッドの上書き
func MethodOverride(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if c.Request().Method == "POST" {
			method := c.Request().PostFormValue("_method")
			if method == "PUT" || method == "DELETE" {
				c.Request().Method = method
			}
		}
		return next(c) //nextとは?
	}
}

// ユーザー全取得
func GetUsers(c echo.Context) error {
	users := []model.User{}
	model.DB.Find(&users)

	return c.JSON(http.StatusOK, users)
}

// ユーザー取得
func GetUser(c echo.Context) error {
	i, _ := strconv.Atoi(c.Param("id"))
	id := uint(i)
	user := model.User{}
	user.FirstById(id)

	return c.JSON(http.StatusOK, user)
}

// ユーザー作成
func CreateUser(c echo.Context) error {
	name := c.FormValue("name")
	p := c.FormValue("password")
	hashed, _ := bcrypt.GenerateFromPassword([]byte(p), 12)
	password := string(hashed)
	belongs := c.FormValue("belongs")
	skills := c.FormValue("skills")

	user := model.User{
		Name:     name,
		Password: password,
		Belongs:  belongs,
		Skills:   skills,
	}
	user.Create()

	return c.JSON(http.StatusOK,user)
}

// ユーザー編集
func UpdateUser(c echo.Context) error {
	i, _ := strconv.Atoi(c.Param("id"))
	id := uint(i)
	name := c.FormValue("name")
	p := c.FormValue("password")
	hashed, _ := bcrypt.GenerateFromPassword([]byte(p), 12)
	password := string(hashed)
	belongs := c.FormValue("belongs")
	skills := c.FormValue("skills")

	user := model.User{
		ID:       id,
		Name:     name,
		Password: password,
		Belongs:  belongs,
		Skills:   skills,
	}
	user.Updates()

	return c.JSON(http.StatusFound,user)
}

// ユーザー削除
func DeleteUser(c echo.Context) error {
	i, _ := strconv.Atoi(c.Param("id"))
	id := uint(i)
	user := model.User{}
	user.DeleteById(id)

	return c.JSON(http.StatusFound, user)
}
