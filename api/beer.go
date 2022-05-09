package api

import (
	"beer-recommend-api/model"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"gopkg.in/gorp.v1"
)

// GetBeers
func GetBeers() echo.HandlerFunc {
	return func(c echo.Context) error {
		id_str := c.Param("id")
		tx := c.Get("Tx").(*gorp.Transaction)

		id, _ := strconv.Atoi(id_str)
		s, err := model.GetBeers(tx, id)
		if err != nil {
			c.Logger().Error("ビール情報取得失敗", err.Error())
			return c.JSON(http.StatusInternalServerError, err)
		}

		return c.JSON(http.StatusOK, s)
	}
}
