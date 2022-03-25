package api

import (
	"beer-recommend-api/model"
	"net/http"

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
		tx := c.Get("Tx").(*gorp.Transaction)

		var l LoginParam
		err := c.Bind(&l)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}

		u, err := model.GethUser(tx, l.UserName)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, err)
		}

		password := []byte(l.Password)
		hashed, err := bcrypt.GenerateFromPassword(password, 10)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}

		err = bcrypt.CompareHashAndPassword(hashed, []byte(l.Password))
		if err == bcrypt.ErrMismatchedHashAndPassword {
			return c.JSON(http.StatusUnauthorized, err)
		}
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}

		return c.JSON(http.StatusOK, u)
	}
}
