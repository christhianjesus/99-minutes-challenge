package entities

import "github.com/golang-jwt/jwt"

type AuthUser struct {
	Username string
	Password string
}

type JwtCustomClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}
