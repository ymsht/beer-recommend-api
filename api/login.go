package api

import (
	"beer-recommend-api/model"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/gorp.v1"
)

const SECRET = "/HeVnSSwSDFI/W8v+YrGOdpXbvpmSARHSRdH4uOW73heR5LqbNAgUw=="

type LoginParam struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

// Login ログインします
func Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		tx := c.Get("Tx").(*gorp.Transaction)

		var l LoginParam
		err := c.Bind(&l)
		if err != nil {
			c.Logger().Error("ログインパラメータバインド失敗", err)
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		u, err := model.GethUser(tx, l.UserName)
		if err != nil {
			c.Logger().Error("ユーザ情報取得失敗", err)
			return echo.ErrUnauthorized
		}

		err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(l.Password))
		if err == bcrypt.ErrMismatchedHashAndPassword {
			c.Logger().Error("パスワード不一致", err)
			return echo.ErrUnauthorized
		}
		if err != nil {
			c.Logger().Error("パスワード比較失敗", err)
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		token := jwt.New(jwt.SigningMethodHS256)
		claims := token.Claims.(jwt.MapClaims)
		claims["name"] = u.UserName
		claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
		tokenString, err := token.SignedString([]byte(SECRET))
		if err != nil {
			c.Logger().Error("トークン変換失敗", err)
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, map[string]string{
			"token": tokenString,
			"name":  u.UserName,
		})
	}
}
