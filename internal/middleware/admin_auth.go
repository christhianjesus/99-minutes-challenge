package middleware

import (
	"interview/domain/constants"
	"net/http"

	"github.com/labstack/echo/v4"
)

func AdminAuth(adminToken string) echo.MiddlewareFunc {
	return func(handler echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			requestToken := c.Request().Header.Get(constants.AdminTokenHeader)

			switch requestToken {
			case adminToken:
				return handler(c)
			case "":
				return c.NoContent(http.StatusBadRequest)
			default:
				return c.NoContent(http.StatusUnauthorized)
			}
		}
	}
}
