package handler

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/middleware"
)

var SECRET = []byte("/HeVnSSwSDFI/W8v+YrGOdpXbvpmSARHSRdH4uOW73heR5LqbNAgUw==")

type JwtCustomClaims struct {
	UID  int    `json:uid`
	Name string `json:"name"`
	jwt.StandardClaims
}

var Config = middleware.JWTConfig{
	Claims:     &JwtCustomClaims{},
	SigningKey: SECRET,
}
