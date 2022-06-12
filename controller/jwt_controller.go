package controller

import (
	"time"
	"net/http"

	"study_go_member/model"

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
	u := new(model.User)
	if err := c.Bind(&u); err != nil {
		return err
	}
	user:=model.User{}
	model.DB.Where("name=?",u.Name).First(&user)

	err:=bcrypt.CompareHashAndPassword([]byte(user.Password),[]byte(u.Password))
	if err!=nil{
		return c.JSON(http.StatusOK,err)
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
		"id": user.ID,
	})
}

//
func Restricted(c echo.Context) error {
	user:=c.Get("user").(*jwt.Token)
	claims:=user.Claims.(*jwtCustomClaims)
	name:=claims.Name
	return c.String(http.StatusOK,"welcome"+name)
}