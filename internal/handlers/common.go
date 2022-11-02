package handlers

import (
	"interview/domain/entities"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type Handler interface {
	RegisterRoutes(*echo.Group, map[string]echo.MiddlewareFunc)
}

func GetUserID(c echo.Context) string {
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(*entities.JwtCustomClaims)

	return claims.Username
}
