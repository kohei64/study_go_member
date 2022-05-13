package controller

import (
	"time"
	"net/http"

	"go-member/model"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"golang.org/x/crypto/bcrypt"
)

type jwtCustomClaims struct{
	UID uint `json:"uid"`
	Name string `json:"name"`
	jwt.StandardClaims
}

var signingkey=[]byte("secret")

var Config=middleware.JWTConfig{
	Claims: &jwtCustomClaims{},
	SigningKey: signingkey,
}

func Login(c echo.Context) error {
	name:=c.FormValue("name")
	password:=c.FormValue("password")

	user:=model.User{}
	model.DB.Where("name=?",name).First(&user)

	err:=bcrypt.CompareHashAndPassword([]byte(user.Password),[]byte(password))
	if err!=nil{
		return echo.ErrUnauthorized
	}

	claims:=&jwtCustomClaims{
		user.ID,
		user.Name,
		jwt.StandardClaims{
			ExpiresAt:time.Now().Add(time.Hour*72).Unix(),
		},
	}

	token:=jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t,err:=token.SignedString(signingkey)
	if err!=nil {
		return err
	}

	return c.JSON(http.StatusOK,echo.Map{
		"token":t,
	})
}

func Restricted(c echo.Context) error {
	user:=c.Get("user").(*jwt.Token)
	claims:=user.Claims.(*jwtCustomClaims)
	name:=claims.Name
	return c.JSON(http.StatusOK,echo.Map{
		"name":name,
	})
}