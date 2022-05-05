package api

import (
	"beer-recommend-api/model"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"gopkg.in/gorp.v1"
)

// GetReviews レビュー情報を返します
func GetReviews() echo.HandlerFunc {
	return func(c echo.Context) error {
		tx := c.Get("Tx").(*gorp.Transaction)
		r, err := model.GetReviews(tx)
		if err != nil {
			c.Logger().Error(err.Error())
			return c.JSON(http.StatusInternalServerError, err)
		}

		return c.JSON(http.StatusOK, r)
	}
}

func GetReviewsSummary() echo.HandlerFunc {
	return func(c echo.Context) error {
		tx := c.Get("Tx").(*gorp.Transaction)
		r, err := model.GetReviewsSummary(tx)
		if err != nil {
			c.Logger().Error(err.Error())
			return c.JSON(http.StatusInternalServerError, err)
		}

		return c.JSON(http.StatusOK, r)
	}
}

func GetReview() echo.HandlerFunc {
	return func(c echo.Context) error {
		id_str := c.Param("id")
		tx := c.Get("Tx").(*gorp.Transaction)

		id, _ := strconv.Atoi(id_str)
		r, err := model.GetReview(tx, id)
		if err != nil {
			c.Logger().Error(err.Error())
			return c.JSON(http.StatusInternalServerError, err)
		}

		return c.JSON(http.StatusOK, r)
	}
}

func CreateReview() echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Get("user").(*jwt.Token)
		claims := token.Claims.(jwt.MapClaims)
		name := claims["name"].(string)

		var r model.Review
		err := c.Bind(&r)
		if err != nil {
			c.Logger().Error(err.Error())
			return c.JSON(http.StatusInternalServerError, err)
		}

		tx := c.Get("Tx").(*gorp.Transaction)

		u, err := model.GethUser(tx, name)
		if err != nil {
			c.Logger().Error(err.Error())
			return c.JSON(http.StatusInternalServerError, err)
		}
		r.MemberId = u.UserId

		err = model.CreateReview(tx, r)
		if err != nil {
			c.Logger().Error(err.Error())
			return c.JSON(http.StatusInternalServerError, err)
		}

		return c.JSON(http.StatusCreated, r)
	}
}
