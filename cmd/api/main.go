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
	"sample.api.kasun.com/pkg/models"
	"sample.api.kasun.com/pkg/models/mysql"
)

type config struct {
	port int
	env  string
	dsn  string
}

type application struct {
	config    config
	logger    *log.Logger
	customers interface {
		Insert(name string, age int, country string, items []string) (int, error)
		GetCustomerById(id int64) (*models.Customer, error)
		GetCustomers() ([]models.Customer, error)
	}
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

	app := &application{
		config:    cfg,
		logger:    logger,
		customers: &mysql.CustomerModel{DB: db},
	}

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	logger.Printf("starting %s server on %s", cfg.env, srv.Addr)
	err = srv.ListenAndServe()
	logger.Fatal(err)

}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
