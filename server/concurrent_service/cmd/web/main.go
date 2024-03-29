package main

import (
	"concsvc/internal/repository"
	"database/sql"
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/alexedwards/scs/redisstore"
	"github.com/alexedwards/scs/v2"
	"github.com/gomodule/redigo/redis"
	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

const httpPort = "9000"

func initDB() *sql.DB {
	conn := connectToDB()
	if conn == nil {
		log.Panic("Can't connect to database")
	}
	return conn
}

func connectToDB() *sql.DB {
	counts := 0

	dsn := os.Getenv("DSN")

	fmt.Println(dsn)

	for {
		connection, err := openDB(dsn)
		if err != nil {
			log.Println("postgres not yet available...")

		} else {
			log.Println("connected to database")
			return connection
		}

		if counts > 10 {
			return nil
		}
		counts++
		log.Println("backing off for 1 second")
		time.Sleep(time.Second * 1)
		continue
	}

}

func openDB(dsn string) (*sql.DB, error) {

	db, err := sql.Open("pgx", dsn)

	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil

}

func initSession() *scs.SessionManager {
	gob.Register(repository.User{})

	session := scs.New()
	session.Store = redisstore.New(initRedis())
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = true

	return session
}

func initRedis() *redis.Pool {
	redisPool := &redis.Pool{
		MaxIdle: 10,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", os.Getenv("REDIS"))
		},
	}
	return redisPool
}

func (app *Config) serve() {
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", httpPort),
		Handler: app.routes(),
	}

	app.InfoLog.Printf("Starting http server on port %s\n", httpPort)

	err := srv.ListenAndServe()

	if err != nil && err != http.ErrServerClosed {
		log.Panic(err)
	}

}

func (app *Config) listenForShutdown() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	app.shutdown()
	os.Exit(0)
}

func (app *Config) shutdown() {
	// perform any cleanup tasks
	app.InfoLog.Println("would run cleanup tasks...")

	// block until waitgroup is empty
	app.Wait.Wait()

	app.Mailer.DoneChan <- true
	app.ErrorChanDone <- true

	app.InfoLog.Println("closing channels and shutting down application...")
	close(app.Mailer.MailerChan)
	close(app.Mailer.ErrorChan)
	close(app.Mailer.DoneChan)
	close(app.ErrorChan)
	close(app.ErrorChanDone)
}

func (app *Config) createMail() Mail {
	errorChan := make(chan error)
	mailerChan := make(chan Message, 128)
	doneChan := make(chan bool)

	m := Mail{
		Domain:      "localhost",
		Host:        "localhost",
		Port:        1025,
		Encryption:  "none",
		FromName:    "admin",
		FromAddress: "info@example.com",
		ErrorChan:   errorChan,
		MailerChan:  mailerChan,
		DoneChan:    doneChan,
		Wait:        app.Wait,
	}

	return m

}

func (app *Config) listenForErrors() {
	for {
		select {
		case err := <-app.ErrorChan:
			app.ErrorLog.Println(err)
		case <-app.ErrorChanDone:
			return
		}
	}
}

func main() {

	db := initDB()

	session := initSession()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	wg := sync.WaitGroup{}

	app := Config{
		Session:       session,
		DB:            db,
		Wait:          &wg,
		InfoLog:       infoLog,
		ErrorLog:      errorLog,
		Models:        repository.New(db),
		ErrorChan:     make(chan error),
		ErrorChanDone: make(chan bool),
	}

	app.Mailer = app.createMail()

	go app.listenForMail()
	go app.listenForShutdown()
	go app.listenForErrors()

	app.serve()
}
