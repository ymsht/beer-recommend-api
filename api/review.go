package api

import (
	"beer-recommend-api/model"
	"net/http"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"gopkg.in/gorp.v1"
)

type Response struct {
	Reviews interface{} `json:"reviews"`
	Total   int         `json:"total"`
}

func GetReviewsSummary() echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Logger().Error("レビュー概要取得")
		tx := c.Get("Tx").(*gorp.Transaction)
		r, err := model.GetReviewsSummary(tx)
		if err != nil {
			c.Logger().Error("レビュー概要取得失敗", err.Error())
			return c.JSON(http.StatusInternalServerError, err)
		}

		var res = Response{r, len(r)}
		return c.JSON(http.StatusOK, res)
	}
}

func GetReview() echo.HandlerFunc {
	return func(c echo.Context) error {
		id_str := c.Param("id")
		tx := c.Get("Tx").(*gorp.Transaction)

		id, _ := strconv.Atoi(id_str)
		r, err := model.GetReview(tx, id)
		if err != nil {
			c.Logger().Error("レビュー取得失敗", err.Error())
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
			c.Logger().Error("モデルパラメータバインド失敗", err.Error())
			return c.JSON(http.StatusInternalServerError, err)
		}

		tx := c.Get("Tx").(*gorp.Transaction)

		u, err := model.GethUser(tx, name)
		if err != nil {
			c.Logger().Error(err.Error())
			return c.JSON(http.StatusInternalServerError, err)
		}
		r.MemberId = u.UserId
		t := time.Now()
		r.Create_date = t
		r.Update_date = t

		err = model.CreateReview(tx, r)
		if err != nil {
			c.Logger().Error("レビュー登録失敗", err.Error())
			return c.JSON(http.StatusInternalServerError, err)
		}

		return c.JSON(http.StatusCreated, r)
	}
}

func DeleteReview() echo.HandlerFunc {
	return func(c echo.Context) error {
		id_str := c.Param("id")
		tx := c.Get("Tx").(*gorp.Transaction)
		id, _ := strconv.Atoi(id_str)

		var r model.Review
		r.ReviewId = id

		cnt, err := model.DeleteReview(tx, r)
		if err != nil {
			c.Logger().Error("レビュー削除失敗", err.Error())
			return c.JSON(http.StatusInternalServerError, err)
		}

		return c.JSON(http.StatusNoContent, cnt)
	}
}

func UpdateReview() echo.HandlerFunc {
	return func(c echo.Context) error {
		id_str := c.Param("id")
		tx := c.Get("Tx").(*gorp.Transaction)
		id, _ := strconv.Atoi(id_str)

		var r model.Review
		err := c.Bind(&r)
		if err != nil {
			c.Logger().Error("モデルパラメータバインド失敗", err.Error())
			return c.JSON(http.StatusInternalServerError, err)
		}
		r.ReviewId = id

		cnt, err := model.UpdateReview(tx, r)
		if err != nil {
			c.Logger().Error("レビュー更新失敗", err.Error())
			return c.JSON(http.StatusInternalServerError, err)
		}

		c.Logger().Info("レビュー更新成功 件数=", cnt)

		return c.JSON(http.StatusNoContent, cnt)
	}
}
