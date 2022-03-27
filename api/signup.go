package api

import (
	"beer-recommend-api/model"
	"net/http"

	"github.com/labstack/echo"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/gorp.v1"
)

type SignupParam struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

// Signup 会員登録します
func Signup() echo.HandlerFunc {
	return func(c echo.Context) error {
		tx := c.Get("Tx").(*gorp.Transaction)

		var s SignupParam
		err := c.Bind(&s)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		hashed, err := bcrypt.GenerateFromPassword([]byte(s.Password), bcrypt.DefaultCost)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		var u model.User
		u.UserName = s.UserName
		u.Password = string(hashed)
		err = model.CreateUser(tx, u)
		if err != nil {
			c.Logger().Error("ユーザ登録失敗:", err)
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, u)
	}
}
