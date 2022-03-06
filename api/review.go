package api

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/ymsht/beer-recommend-api/model"
	"gopkg.in/gorp.v1"
)

// GetAlongs 沿線情報を返します
func GetReviews() echo.HandlerFunc {
	return func(c echo.Context) error {
		tx := c.Get("Tx").(*gorp.Transaction)

		alongs, err := model.GetReviews(tx)
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, alongs)
	}
}
