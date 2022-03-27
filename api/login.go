package api

import (
	"beer-recommend-api/model"
	"net/http"

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
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		u, err := model.GethUser(tx, l.UserName)
		if err != nil {
			return echo.ErrUnauthorized
		}

		// 会員登録で使う処理
		// hashed, err := bcrypt.GenerateFromPassword([]byte(l.Password), bcrypt.DefaultCost)
		// if err != nil {
		// 	return c.JSON(http.StatusInternalServerError, err)
		// }
		// c.Logger().Error("hash:", string(hashed))

		err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(l.Password))
		if err == bcrypt.ErrMismatchedHashAndPassword {
			c.Logger().Error("パスワード不一致")
			return echo.ErrUnauthorized
		}
		if err != nil {
			c.Logger().Error(err)
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		token := jwt.New(jwt.SigningMethodHS256)
		claims := token.Claims.(jwt.MapClaims)
		claims["userName"] = u.UserName
		tokenString, _ := token.SignedString([]byte(SECRET))

		return c.JSON(http.StatusOK, map[string]string{"token": tokenString})
	}
}
