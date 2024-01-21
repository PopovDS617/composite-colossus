package api

import (
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

func TestAuthenticateSuccess(t *testing.T) {
	testDB := setupDB(t)
	app := utils.SetupFiber()

	preparedUser, _ := types.NewUserFromParams(testUser)

	_, err := testDB.UserStore.Insert(context.Background(), preparedUser)

	if err != nil {
		t.Error(err)
	}

	defer testDB.teardown(t)

	authHandler := NewAuthHandler(testDB.UserStore)

	app.Post("/auth", authHandler.HandleAuth)

	authParams := AuthParams{
		Email:    testUser.Email,
		Password: testUser.Password,
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

	preparedUser.EncryptedPassword = ""

	if !reflect.DeepEqual(preparedUser, resUser) {
		t.Fatalf("expected the response user to be equal to the inserted user")
	}

}

func TestAuthenticateFailure(t *testing.T) {
	testDB := setupDB(t)
	app := utils.SetupFiber()

	preparedUser, _ := types.NewUserFromParams(testUser)

	_, err := testDB.UserStore.Insert(context.Background(), preparedUser)

	if err != nil {
		t.Error(err)
	}

	defer testDB.teardown(t)

	authHandler := NewAuthHandler(testDB.UserStore)

	app.Post("/auth", authHandler.HandleAuth)

	authParams := AuthParams{
		Email:    testUser.Email + "fail",
		Password: testUser.Password,
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

	var response GenericResponse

	utils.DecodeResBody[*GenericResponse](res, &response)

	resType := response.Type
	resMessage := response.Message

	if resType != "error" {
		t.Fatalf("expected response type to be 'error' but git %s", resType)
	}

	if resMessage != "invalid credentials" {
		t.Fatalf("expected response type to be 'invalid credentials' but git %s", resMessage)
	}

}
