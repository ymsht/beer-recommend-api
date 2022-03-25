package api

import (
	"io/ioutil"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

type JwtCustomClaims struct {
	Name  string `json:"name"`
	Admin bool   `json:"admin"`
	jwt.StandardClaims
}

func Restricted() echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Logger().Error("返却")
		token, ok := c.Get("user").(*jwt.Token)
		if !ok {
			return c.JSON(http.StatusOK, "失敗2")
		}
		c.Logger().Error(token)

		c.Logger().Error("返却2")
		claims, ok := token.Claims.(jwt.MapClaims)
		if ok && token.Valid {
			return c.JSON(http.StatusOK, claims["member_id"])
		} else {
			return c.JSON(http.StatusOK, "失敗")
		}

		// verifyBytes, err := ioutil.ReadFile("./id_rsa.pub.pkcs8")
		// if err != nil {
		// 	c.Logger().Error(err.Error())
		// 	return c.JSON(http.StatusInternalServerError, err)
		// }
		// verifyKey, err := jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
		// if err != nil {
		// 	c.Logger().Error(err.Error())
		// 	return c.JSON(http.StatusInternalServerError, err)
		// }

		// token, err := request.ParseFromRequest(c.Request(), request.AuthorizationHeaderExtractor, func(token *jwt.Token) (interface{}, error) {
		// 	_, err := token.Method.(*jwt.SigningMethodRSA)
		// 	if !err {
		// 		return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		// 	} else {
		// 		return verifyKey, nil
		// 	}
		// })
		// if err == nil && token.Valid {
		// 	return c.JSON(http.StatusOK, echo.Map{"status": "ok"})
		// } else {
		// 	return c.JSON(http.StatusUnauthorized, echo.Map{"status": "n1"})
		// }

		// token, ok := c.Get("user").(*jwt.Token)
		// if !ok {
		// 	return c.JSON(http.StatusOK, echo.Map{"status": "n1"})
		// }

		// if err == nil && token.Valid {
		// 	return c.JSON(http.StatusOK, echo.Map{"status": "ok"})
		// } else {
		// 	return c.JSON(http.StatusOK, echo.Map{"status": "ng"})
		// }
	}
}

func Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		//tx := c.Get("Tx").(*gorp.Transaction)

		username := c.FormValue("username")
		password := c.FormValue("password")

		username = "jon"
		password = "shhh!"

		// Throws unauthorized error
		if username != "jon" || password != "shhh!" {
			return echo.ErrUnauthorized
		}

		signBytes, err := ioutil.ReadFile("es256.key.pkcs8")
		if err != nil {
			c.Logger().Error("失敗失敗1")
			c.Logger().Error(err.Error())
			return c.JSON(http.StatusInternalServerError, err)
		}

		signKey, err := jwt.ParseECPrivateKeyFromPEM(signBytes)
		if err != nil {
			c.Logger().Error("失敗失敗2")
			c.Logger().Error(err.Error())
			return c.JSON(http.StatusInternalServerError, err)
		}

		token := jwt.New(jwt.SigningMethodHS256)
		claims := token.Claims.(jwt.MapClaims)
		claims["member_id"] = 12345
		claims["admin"] = false
		claims["iat"] = time.Now()
		claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

		tokenString, err := token.SignedString(signKey)
		if err != nil {
			c.Logger().Error(err.Error())
			return c.JSON(http.StatusInternalServerError, err)
		}

		return c.JSON(http.StatusOK, echo.Map{
			"token": tokenString,
		})
	}
}
