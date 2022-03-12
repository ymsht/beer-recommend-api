package api

import (
	"beer-recommend-api/model"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"gopkg.in/gorp.v1"
)

// GetReviews レビュー情報を返します
func GetReviews() echo.HandlerFunc {
	return func(c echo.Context) error {
		tx := c.Get("Tx").(*gorp.Transaction)

		reviews, err := model.GetReviews(tx)
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, reviews)
	}
}

func GetReview() echo.HandlerFunc {
	return func(c echo.Context) error {
		id_str := c.Param("id")
		tx := c.Get("Tx").(*gorp.Transaction)

		id, _ := strconv.Atoi(id_str)
		review, err := model.GetReview(tx, id)
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, review)
	}
}

func CreateReview() echo.HandlerFunc {
	return func(c echo.Context) error {
		var r model.Review
		err := c.Bind(&r)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}

		tx := c.Get("Tx").(*gorp.Transaction)

		err = model.CreateReview(tx, r)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}

		return c.JSON(http.StatusCreated, r)
	}
}
