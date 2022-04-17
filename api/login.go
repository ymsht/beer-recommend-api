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
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		u, err := model.GethUser(tx, l.UserName)
		if err != nil {
			return echo.ErrUnauthorized
		}

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
		claims["name"] = u.UserName
		claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
		tokenString, _ := token.SignedString([]byte(SECRET))

		return c.JSON(http.StatusOK, map[string]string{"token": tokenString})
	}
}
