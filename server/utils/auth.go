package utils

import (
	"app/api/custerr"
	"app/types"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func ValidateToken(inputToken string) (jwt.MapClaims, error) {

	token, err := jwt.Parse(inputToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			fmt.Println("invalid signing method")
			return nil, custerr.Unauthorized()
		}

		secret := os.Getenv("JWT_SECRET")

		return []byte(secret), nil
	})

	if err != nil {
		return nil, custerr.Unauthorized()
	}

	if !token.Valid {
		fmt.Println("invalid token")
		return nil, custerr.Unauthorized()
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, custerr.Unauthorized()
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

	secret := os.Getenv("JWT_SECRET")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims, nil)

	tokenStr, err := token.SignedString([]byte(secret))

	if err != nil {

		log.Fatal(err)
	}

	return tokenStr

}

func GetUserFromContext(ctx *fiber.Ctx) (*types.User, error) {

	user, ok := ctx.Context().UserValue("user").(*types.User)

	if !ok {
		return nil, custerr.Unauthorized()
	}

	return user, nil
}
