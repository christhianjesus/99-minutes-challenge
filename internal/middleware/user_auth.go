package middleware

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

const (
	UserAuthHeader = "auth_user"
)

func UserAuth() echo.MiddlewareFunc {
	return func(handler echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authUser := c.Request().Header.Get(UserAuthHeader)
			if authUser == "" {
				return c.NoContent(http.StatusUnauthorized)
			}
			return handler(c)
		}
	}
}

func GetUserID(c echo.Context) string {
	return c.Request().Header.Get(UserAuthHeader)
}
