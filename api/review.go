package api

import (
	"beer-recommend-api/model"
	"net/http"

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
		//id := c.Param("review_id")
		tx := c.Get("Tx").(*gorp.Transaction)

		review, err := model.GetReview(tx)
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
		return c.JSON(http.StatusOK, r)

		// tx := c.Get("Tx").(*gorp.Transaction)

		// err := model.CreateReview(tx, r)
		// if err != nil {
		// 	return c.JSON(http.StatusInternalServerError, err)
		// }

		// return c.JSON(http.StatusOK, "ok")
	}
}
