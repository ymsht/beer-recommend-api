package api

import (
	"beer-recommend-api/handler"
	"beer-recommend-api/model"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/gorp.v1"
)

type LoginParam struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

// Login ログインします
func Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Logger().Error("ログイン")

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

		claims := &handler.JwtCustomClaims{
			u.UserId,
			u.UserName,
			jwt.StandardClaims{
				ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
			},
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		t, err := token.SignedString(handler.SECRET)
		if err != nil {
			c.Logger().Error("トークン変換失敗", err)
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		// c.Logger().Error("JWT トークン取得")
		// user := c.Get("user").(*jwt.Token)
		// c.Logger().Error("JWT クレーム取得")
		// cl := user.Claims.(jwt.MapClaims)
		// c.Logger().Error("JWT ユーザ名取得")
		// name := cl["name"].(string)
		// c.Logger().Error("JWT ユーザ名", name)

		return c.JSON(http.StatusOK, map[string]string{
			"token": t,
			"name":  u.UserName,
		})
	}
}
