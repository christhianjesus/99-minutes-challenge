package handlers

import (
	"interview/domain/constants"
	"interview/domain/entities"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	secret []byte
}

func NewAuthHandler(secret string) Handler {
	return &AuthHandler{[]byte(secret)}
}

func (h *AuthHandler) RegisterRoutes(router *echo.Group, _ map[string]echo.MiddlewareFunc) {
	router.POST(constants.AuthPath, h.Login)
}

func (h *AuthHandler) Login(c echo.Context) error {
	auth := &entities.AuthUser{}
	if err := c.Bind(auth); err != nil {
		return err
	}

	if constants.AuthUsers[auth.Username] != auth.Password {
		return echo.ErrUnauthorized
	}

	// Set custom claims
	claims := &entities.JwtCustomClaims{
		Username: auth.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour).Unix(),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString(h.secret)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token": signedToken,
	})
}
