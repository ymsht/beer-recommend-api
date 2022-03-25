package router

import (
	"beer-recommend-api/api"
	"beer-recommend-api/db"
	"beer-recommend-api/handler"
	mw "beer-recommend-api/middleware"
	"io/ioutil"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// Init 初期化
func Init() *echo.Echo {
	e := echo.New()

	fp, err := os.OpenFile("/var/log/echo.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	e.Use(middleware.Logger())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Output: fp,
	}))

	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET},
		AllowHeaders: []string{
			"Access-Control-Allow-Headers",
			"Content-Type", "Content-Length",
			"Accept-Encoding",
			"X-CSRF-Token",
			"Authorization"},
	}))
	e.HTTPErrorHandler = handler.CustomHTTPErrorHandler

	e.Use(mw.TransactionHandler(db.Init()))

	e.POST("/api/v1/login", api.Login())

	keyData, err := ioutil.ReadFile("es256.pub.key")
	if err != nil {
		panic(err)
	}
	key, err := jwt.ParseECPublicKeyFromPEM(keyData)
	if err != nil {
		panic(err)
	}

	v1 := e.Group("/api/v1")
	{
		config := middleware.JWTConfig{
			Claims:     &api.JwtCustomClaims{},
			SigningKey: key,
		}
		v1.Use(middleware.JWTWithConfig(config))

		v1.GET("/reviews", api.GetReviews())
		v1.GET("/review/:id", api.GetReview())
		v1.POST("/review", api.CreateReview())

		v1.GET("/flavors", api.GetFlavors())

		v1.GET("/styles", api.GetStyles())

		v1.GET("/countries", api.GetCountries())

		v1.GET("/areas", api.GetAreas())

		v1.GET("/restricted", api.Restricted())
	}

	return e
}
