package api

import (
	"beer-recommend-api/model"
	"net/http"

	"github.com/labstack/echo"
	"gopkg.in/gorp.v1"
)

// GetStyles スタイル情報を返します
func GetStyles() echo.HandlerFunc {
	return func(c echo.Context) error {
		tx := c.Get("Tx").(*gorp.Transaction)

		s, err := model.GetStyles(tx)
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, s)
	}
}
