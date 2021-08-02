package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"sample.api.kasun.com/cmd/api/helpers"
	"sample.api.kasun.com/cmd/api/repository/mysql"
	"sample.api.kasun.com/cmd/api/service"
)

type config struct {
	port int
	env  string
	dsn  string
}

type application struct {
	config          config
	logger          *log.Logger
	customerService service.CustomerService
}

func main() {

	cfg := config{
		port: 4000,
		env:  "development",
		dsn:  "",
	}

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	flag.StringVar(&cfg.dsn, "db-dsn", os.Getenv("CUSTOMER_API_DB_DSN"), "MYSQL DSN")

	db, err := sql.Open("mysql", cfg.dsn)
	if err != nil {
		logger.Fatal(err)
	}

	defer db.Close()

	rp := &mysql.CustomerModel{DB: db}

	cs := service.DefaultCustomerService{
		Repository: rp,
		Logger:     logger,
		Helpers:    helpers.Helpers{},
		Errors:     helpers.Errors{},
	}

	app := &application{
		config:          cfg,
		logger:          logger,
		customerService: cs,
	}

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	go app.consumeKafka()

	logger.Printf("starting %s server on %s", cfg.env, srv.Addr)
	err = srv.ListenAndServe()
	logger.Fatal(err)

}
