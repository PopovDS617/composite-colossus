package types

import (
	"fmt"
	"regexp"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

const (
	bcryptCost      = 12
	minFirstNameLen = 2
	minLastNameLen  = 2
	minPasswordLen  = 6
)

type User struct {
	ID                primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"` // or json:"_"
	FirstName         string             `bson:"first_name" json:"first_name"`
	LastName          string             `bson:"last_name" json:"last_name"`
	Email             string             `bson:"email" json:"email"`
	EncryptedPassword string             `bson:"encrypted_password" json:"-"`
}

type CreateUserParams struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type UpdateUserParams struct {
	FirstName string `bson:"first_name" json:"first_name"`
	LastName  string `bson:"last_name" json:"last_name"`
}

func NewUserFromParams(params CreateUserParams) (*User, error) {

	encpw, err := bcrypt.GenerateFromPassword([]byte(params.Password), bcryptCost)

	if err != nil {
		return nil, err
	}

	return &User{FirstName: params.FirstName, LastName: params.LastName, Email: params.Email, EncryptedPassword: string(encpw)}, nil

}

func (params CreateUserParams) ValidateUser() map[string]string {

	errorsMap := map[string]string{}

	if len(params.FirstName) < minFirstNameLen {
		errorsMap["first_name"] = fmt.Sprintf("length should be at least %d characters", minFirstNameLen)
	}
	if len(params.LastName) < minLastNameLen {
		errorsMap["last_name"] = fmt.Sprintf("length should be at least %d characters", minLastNameLen)
	}
	if len(params.Password) < minPasswordLen {
		errorsMap["password"] = fmt.Sprintf("length should be at least %d characters", minPasswordLen)
	}
	if !IsEmailValid(params.Email) {
		errorsMap["email"] = "is invalid"
	}

	return errorsMap
}

func (p UpdateUserParams) ToBSON() bson.M {

	m := bson.M{}

	if len(p.FirstName) > 0 {
		m["first_name"] = p.FirstName
	}

	if len(p.LastName) > 0 {
		m["last_name"] = p.LastName
	}

	return m
}

func IsEmailValid(e string) bool {
	emailRegex := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	return emailRegex.MatchString(e)
}
