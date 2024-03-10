package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (app *Config) routes() http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(app.SessionLoad)

	mux.Get("/", app.GetHomePage)
	mux.Get("/login", app.GetLoginPage)
	mux.Post("/login", app.PostLoginPage)
	mux.Get("/logout", app.Logout)
	mux.Get("/register", app.GetRegisterPage)
	mux.Post("/register", app.PostRegisterPage)
	mux.Post("/activate", app.ActivateAccount)
	mux.Get("/test-email", func(w http.ResponseWriter, r *http.Request) {

		m := Mail{
			Domain:      "localhost",
			Host:        "localhost",
			Port:        1025,
			Encryption:  "none",
			FromName:    "admin",
			FromAddress: "info@example.com",
			ErrorChan:   make(chan error),
		}

		msg := Message{
			To:      "me@example.com",
			Subject: "Test",
			Data:    "Hello there",
		}

		m.sendMail(msg, make(chan error))
	})

	return mux
}
