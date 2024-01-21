package utils

import (
	"app/types"
	"fmt"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func ValidateToken(inputToken string) (jwt.MapClaims, error) {

	token, err := jwt.Parse(inputToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			fmt.Println("invalid signing method")
			return nil, fmt.Errorf("unauthorized")
		}

		// secret := os.Getenv("JWT_SECRET")

		// TODO: add proper env
		secret := []byte("secretsupersecret")

		return []byte(secret), nil
	})
	if err != nil {
		return nil, fmt.Errorf("unathorized")
	}

	if !token.Valid {
		fmt.Println("invalid token")
		return nil, fmt.Errorf("unathorized")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("unathorized")
	}

	return claims, nil

}

func IsPasswordValid(receivedPw, savedPw string) bool {
	return bcrypt.CompareHashAndPassword([]byte(savedPw), []byte(receivedPw)) == nil
}

func CreateToken(user *types.User) string {

	createdAt := time.Now()
	expires := createdAt.Add(time.Hour * 6).Unix()

	claims := jwt.MapClaims{
		"user_id": user.ID,
		"email":   user.Email,
		"expires": expires,
	}

	// secret := os.Getenv("JWT_SECRET")

	// TODO: add proper env
	secret := []byte("secretsupersecret")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims, nil)

	tokenStr, err := token.SignedString(secret)

	if err != nil {

		log.Fatal(err)
	}

	return tokenStr

}
