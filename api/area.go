package api

import (
	"beer-recommend-api/model"
	"net/http"

	"github.com/labstack/echo"
	"gopkg.in/gorp.v1"
)

// GetAreas 地域情報を返します
func GetAreas() echo.HandlerFunc {
	return func(c echo.Context) error {
		tx := c.Get("Tx").(*gorp.Transaction)

		a, err := model.GetAreas(tx)
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, a)
	}
}
