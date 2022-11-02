package middleware

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

const (
	AdminTokenHeader = "AdminToken"
)

func AdminAuth(adminToken string) echo.MiddlewareFunc {
	return func(handler echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			requestToken := c.Request().Header.Get(AdminTokenHeader)
			if requestToken != adminToken {
				return c.NoContent(http.StatusUnauthorized)
			}
			return handler(c)
		}
	}
}
