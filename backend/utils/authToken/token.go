package authToken

import (
	"fmt"
	"log"
	"os"
	"sinarmas/kredit-sinarmas/models"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const ExpTimeMinute = 60

func GenerateToken(user models.User) (string, error) {
	actClaims := Claims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * ExpTimeMinute).Unix(),
		},
		ID:       user.ID,
		Username: user.Username,
		Name:     user.Name,
	}

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, actClaims)
	resultToken, err := at.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		log.Println("Login: Error while creating token")
		return "", fmt.Errorf("Error while creating token")
	}

	return resultToken, nil
}

// func ValidationToken(dataTime int64) bool {
// 	if dataTime > time.Now().Unix() {
// 		log.Println("Token Validation: Token Not Expired")
// 		return false
// 	}

// 	log.Println("Token Validation: Expired")
// 	return true
// }
