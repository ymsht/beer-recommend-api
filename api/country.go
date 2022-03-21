package api

import (
	"beer-recommend-api/model"
	"net/http"

	"github.com/labstack/echo"
	"gopkg.in/gorp.v1"
)

// GetCountries 原産国情報を返します
func GetCountries() echo.HandlerFunc {
	return func(c echo.Context) error {
		tx := c.Get("Tx").(*gorp.Transaction)

		countries, err := model.GetCountries(tx)
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, countries)
	}
}
