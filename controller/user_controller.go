package controller

import (
	"net/http"
	"strconv"

	"study_go_member/model"

	"github.com/labstack/echo"
	// "golang.org/x/crypto/bcrypt"
)

// メソッドの上書き
func MethodOverride(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if c.Request().Method == "POST" {
			method := c.Request().PostFormValue("_method")
			if method == "PUT" || method == "PATCH" || method == "DELETE" {
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
	model.DB.Where("id = ?", id).First(&user)

	return c.JSON(http.StatusOK, user)
}

//ユーザー作成
func CreateUser(c echo.Context) error {
	u := new(model.User)
	if err := c.Bind(&u); err != nil {
		return err
	}
	//todo:password hash化する
	//todo:validation追加
	model.DB.Create(&u)

	return c.JSON(http.StatusOK, u)
}

// ユーザー編集
func UpdateUser(c echo.Context) error {
	i, _ := strconv.Atoi(c.Param("id"))
	id := uint(i)
	u := new(model.User)
	if err:=c.Bind(&u); err!=nil{
		return err
	}
	u.ID = id
	model.DB.Updates(&u)

	return c.JSON(http.StatusOK, u)
}

// ユーザー削除
func DeleteUser(c echo.Context) error {
	i,_ := strconv.Atoi(c.Param("id"))
	id := uint(i)
	user := model.User{}
	model.DB.Where("id=?", id).Delete(&user)

	return c.JSON(http.StatusOK, user)
}