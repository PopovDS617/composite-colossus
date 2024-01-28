package api

import (
	"app/db/fixtures"
	"app/types"
	"app/utils"
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

var authHandlerTestUser = types.CreateUserParams{
	FirstName: "Mark",
	LastName:  "One",
	Email:     "Mark@mail.com",
	Password:  "Mark_One",
}

func TestAuthenticateSuccess(t *testing.T) {
	testDB := setupDB(t)
	app := utils.SetupFiber()

	defer testDB.teardown(t)

	insertedUser := fixtures.AddUser(&testDB.Store, authHandlerTestUser)

	authHandler := NewAuthHandler(testDB.User)

	app.Post("/auth", authHandler.HandleAuth)

	authParams := AuthParams{
		Email:    authHandlerTestUser.Email,
		Password: authHandlerTestUser.Password,
	}

	b, _ := json.Marshal(authParams)

	req := httptest.NewRequest("POST", "/auth", bytes.NewReader(b))
	req.Header.Add("Content-Type", "application/json")

	res, err := app.Test(req)

	if err != nil {
		t.Error(err)
	}

	if res.StatusCode != http.StatusOK {
		t.Errorf("expected http status of 200 but got %d", res.StatusCode)
	}

	var response AuthResponse

	utils.DecodeResBody[*AuthResponse](res, &response)

	resUser := response.User
	resToken := response.Token

	if resToken == "" {
		t.Fatalf("token is not present")
	}

	insertedUser.EncryptedPassword = ""

	if !reflect.DeepEqual(insertedUser, resUser) {
		t.Fatalf("expected the response user to be equal to the inserted user")
	}

}

func TestAuthenticateFailure(t *testing.T) {
	testDB := setupDB(t)
	app := utils.SetupFiber()

	preparedUser, _ := types.NewUserFromParams(authHandlerTestUser)

	_, err := testDB.User.Insert(context.Background(), preparedUser)

	if err != nil {
		t.Error(err)
	}

	defer testDB.teardown(t)

	authHandler := NewAuthHandler(testDB.User)

	app.Post("/auth", authHandler.HandleAuth)

	authParams := AuthParams{
		Email:    authHandlerTestUser.Email + "fail",
		Password: authHandlerTestUser.Password,
	}

	b, _ := json.Marshal(authParams)

	req := httptest.NewRequest("POST", "/auth", bytes.NewReader(b))
	req.Header.Add("Content-Type", "application/json")

	res, err := app.Test(req)

	if err != nil {
		t.Error(err)
	}

	if res.StatusCode != http.StatusBadRequest {
		t.Errorf("expected http status of 400 but got %d", res.StatusCode)
	}

}
