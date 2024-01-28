package api

import (
	"app/types"
	"app/utils"
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"
)

var userHandlerTestUser = types.CreateUserParams{
	FirstName: "Mark",
	LastName:  "One",
	Email:     "Mark@mail.com",
	Password:  "Mark_One",
}

func TestPostUser(t *testing.T) {
	testDB := setupDB(t)
	app := utils.SetupFiber()

	defer testDB.teardown(t)

	userHandler := NewUserHandler(testDB.User)

	app.Post("/users", userHandler.HandlePostUser)

	b, _ := json.Marshal(userHandlerTestUser)

	req := httptest.NewRequest("POST", "/users", bytes.NewReader(b))
	req.Header.Add("Content-Type", "application/json")

	res, _ := app.Test(req)

	var resUser types.User

	utils.DecodeResBody[*types.User](res, &resUser)

	if len(resUser.ID) == 0 {
		t.Errorf("no user id in the response")
	}
	if len(resUser.EncryptedPassword) > 0 {
		t.Errorf("password should not be in the response")
	}

	if resUser.FirstName != userHandlerTestUser.FirstName {
		t.Errorf("expected firstname %s but received %s", userHandlerTestUser.FirstName, resUser.FirstName)
	}
	if resUser.LastName != userHandlerTestUser.LastName {
		t.Errorf("expected firstname %s but received %s", userHandlerTestUser.LastName, resUser.LastName)
	}
	if resUser.Email != userHandlerTestUser.Email {
		t.Errorf("expected firstname %s but received %s", userHandlerTestUser.Email, resUser.Email)
	}
}
