package api

import (
	"app/db"
	"app/types"
	"app/utils"
	"bytes"
	"context"
	"encoding/json"
	"log"
	"net/http/httptest"
	"testing"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type testdb struct {
	db.UserStore
}

const testdburi = "mongodb://root:password@localhost:27017/?authSource=admin"

func setupDB(t *testing.T) *testdb {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(testdburi))

	if err != nil {
		log.Fatal(err)
	}

	return &testdb{
		UserStore: db.NewMongoUserStore(client, "test"),
	}
}

func (tdb *testdb) teardown(t *testing.T) {
	if err := tdb.UserStore.Drop(context.TODO()); err != nil {
		t.Fatal(err)
	}
}

func TestPostUser(t *testing.T) {
	testDB := setupDB(t)
	app := utils.SetupFiber()

	defer testDB.teardown(t)

	userHandler := NewUserHandler(testDB.UserStore)

	app.Post("/users", userHandler.HandlePostUser)

	updUser := types.CreateUserParams{
		Email:     "test@test.com",
		FirstName: "Mike",
		LastName:  "Elson",
		Password:  "password123",
	}

	b, _ := json.Marshal(updUser)

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

	if resUser.FirstName != updUser.FirstName {
		t.Errorf("expected firstname %s but received %s", updUser.FirstName, resUser.FirstName)
	}
	if resUser.LastName != updUser.LastName {
		t.Errorf("expected firstname %s but received %s", updUser.LastName, resUser.LastName)
	}
	if resUser.Email != updUser.Email {
		t.Errorf("expected firstname %s but received %s", updUser.Email, resUser.Email)
	}
}
