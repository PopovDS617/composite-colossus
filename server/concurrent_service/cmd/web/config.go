package main

import (
	"concsvc/internal/repository"
	"database/sql"
	"log"
	"sync"

	"github.com/alexedwards/scs/v2"
)

type Config struct {
	Session       *scs.SessionManager
	DB            *sql.DB
	InfoLog       *log.Logger
	ErrorLog      *log.Logger
	Wait          *sync.WaitGroup
	Models        repository.Models
	Mailer        Mail
	ErrorChan     chan error
	ErrorChanDone chan bool
}
