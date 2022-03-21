package api

import (
	"beer-recommend-api/model"
	"net/http"

	"github.com/labstack/echo"
	"gopkg.in/gorp.v1"
)

// GetFlavors フレーバー情報を返します
func GetFlavors() echo.HandlerFunc {
	return func(c echo.Context) error {
		tx := c.Get("Tx").(*gorp.Transaction)

		f, err := model.GetFlavors(tx)
		if err != nil {
			c.Logger().Error(err.Error())
			return c.JSON(http.StatusInternalServerError, err)
		}

		return c.JSON(http.StatusOK, f)
	}
}
