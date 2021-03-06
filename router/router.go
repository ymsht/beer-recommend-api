package router

import (
	"beer-recommend-api/api"
	"beer-recommend-api/db"
	"beer-recommend-api/handler"
	mw "beer-recommend-api/middleware"
	"os"

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
		AllowMethods: []string{echo.GET, echo.POST, echo.DELETE, echo.PUT},
		AllowHeaders: []string{
			"Access-Control-Allow-Headers",
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
			"X-CSRF-Token",
			"Authorization"},
	}))
	e.HTTPErrorHandler = handler.CustomHTTPErrorHandler

	e.Use(mw.TransactionHandler(db.Init()))

	e.POST("/api/login", api.Login())
	e.POST("/api/signup", api.Signup())

	v1 := e.Group("/api/v1")
	{
		v1.Use(middleware.JWTWithConfig(handler.Config))
		// v1.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		// 	return func(c echo.Context) error {
		// 		if err := next(c); err != nil {
		// 			c.Error(err)
		// 		}

		// 		output := fmt.Sprintf("%#v", c.Request().Header)

		// 		os.Stdout.Write([]byte(output + "\n"))
		// 		return next(c)
		// 	}
		// })

		v1.GET("/reviews/summary", api.GetReviewsSummary())
		v1.GET("/review/:id", api.GetReview())
		v1.POST("/review", api.CreateReview())
		v1.DELETE("/review/:id", api.DeleteReview())
		v1.PUT("/review/:id", api.UpdateReview())

		v1.GET("/flavors", api.GetFlavors())

		v1.GET("/styles", api.GetStyles())

		v1.GET("/countries", api.GetCountries())

		v1.GET("/areas", api.GetAreas())

		v1.GET("/beers/:id", api.GetBeers())

		v1.GET("/breweries", api.GetBreweries())
	}

	return e
}
