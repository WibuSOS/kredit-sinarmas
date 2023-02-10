package authToken

import "github.com/dgrijalva/jwt-go"

type Claims struct {
	jwt.StandardClaims
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Name     string `json:"name"`
}
