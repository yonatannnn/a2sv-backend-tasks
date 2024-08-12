package Infrastructure

import (
	domain "api/task_manager/Domain"
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtSecretKey = []byte(os.Getenv("SECRET_KEY"))

func GenerateJWT(user domain.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"username":   user.Username,
		"role":    user.Role,
		"exp":     time.Now().Add(time.Hour * 24 * 3).Unix(),
	})

	tokenString, err := token.SignedString(jwtSecretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateJWT(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtSecretKey, nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}