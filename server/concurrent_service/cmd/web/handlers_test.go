package main

import (
	"concsvc/internal/repository"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

var pageTests = []struct {
	name               string
	url                string
	expectedStatusCode int
	handler            http.HandlerFunc
	sessionData        map[string]any
	expectedHTML       string
}{
	{
		name:               "Home",
		url:                "/",
		expectedStatusCode: http.StatusOK,
		handler:            testApp.GetHomePage},
	{
		name:               "Login",
		url:                "/login",
		expectedStatusCode: http.StatusOK,
		handler:            testApp.GetLoginPage,
		expectedHTML:       `<h1 class="mt-5">Login</h1>`,
	},
	{
		name:               "Logout",
		url:                "/logout",
		expectedStatusCode: http.StatusSeeOther,
		handler:            testApp.Logout,
		sessionData: map[string]any{
			"userID": 1,
			"user":   repository.User{},
		},
	},
}

func Test_Get_Pages(t *testing.T) {

	for _, e := range pageTests {
		rr := httptest.NewRecorder()

		req, _ := http.NewRequest("GET", e.url, nil)
		ctx := getCtx(req)

		req = req.WithContext(ctx)

		if len(e.sessionData) > 0 {
			for k, v := range e.sessionData {
				testApp.Session.Put(ctx, k, v)
			}
		}

		e.handler.ServeHTTP(rr, req)

		if rr.Code != e.expectedStatusCode {
			t.Errorf("expected response code to be %d, received %d\n", e.expectedStatusCode, rr.Code)
		}

		if len(e.expectedHTML) > 0 {
			html := rr.Body.String()
			if !strings.Contains(html, e.expectedHTML) {
				t.Errorf("expected to find %s in %s", e.expectedHTML, e.name)
			}
		}
	}

}

func Test_Post_Pages(t *testing.T) {

	postedData := url.Values{
		"email":    {"admin@example.com"},
		"password": {"abc123abc123abc123abc123"},
	}

	rr := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/login", strings.NewReader(postedData.Encode()))
	ctx := getCtx(req)
	req = req.WithContext(ctx)

	handler := http.HandlerFunc(testApp.PostLoginPage)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusSeeOther {
		t.Error("wrong code returned")
	}

	if !testApp.Session.Exists(ctx, "userID") {
		t.Error("did not find userID in session")
	}

}

func TestConfig_SubscribeToPlan(t *testing.T) {
	rr := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/subscribe?id=1", nil)
	ctx := getCtx(req)
	req = req.WithContext(ctx)

	testApp.Session.Put(ctx, "user", repository.User{
		ID:        1,
		Email:     "admin@example.com",
		FirstName: "Admin",
		LastName:  "User",
		Active:    1,
	})

	handler := http.HandlerFunc(testApp.SubscribeToPlan)

	handler.ServeHTTP(rr, req)

	testApp.Wait.Wait()

	if rr.Code != http.StatusSeeOther {
		t.Errorf("expected status code of statusseeother, but got %d", rr.Code)
	}

}
