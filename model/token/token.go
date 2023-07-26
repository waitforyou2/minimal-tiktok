package token

import "github.com/golang-jwt/jwt/v5"

type MyCustomClaims struct {
	Name     string `json:"username"`
	Password string `json:"password"`
	jwt.RegisteredClaims
}
