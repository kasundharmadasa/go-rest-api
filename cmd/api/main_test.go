package main

import (
	"log"
	"os"

	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"sample.api.kasun.com/cmd/api/helpers"
	"sample.api.kasun.com/cmd/api/repository/mock"
	"sample.api.kasun.com/cmd/api/service"
)

func TestGetCustomers(t *testing.T) {

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)
	rp := &mock.CustomerModel{}

	cs := service.DefaultCustomerService{
		Repository: rp,
		Logger:     logger,
		Helpers:    helpers.Helpers{},
		Errors:     helpers.Errors{},
	}
	app := &application{
		logger:          logger,
		customerService: cs,
	}

	req, err := http.NewRequest("GET", "/customers", nil)
	if err != nil {
		t.Fatal(err)
	}
	reqRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(app.customerService.GetCustomers)
	handler.ServeHTTP(reqRecorder, req)

	// Test whether the response is OK
	if status := reqRecorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Test the expected response body
	expected := `{"customers":[{"ID":1,"Name":"Bob","Age":22,"Country":"MT",`
	if !strings.Contains(reqRecorder.Body.String(), expected) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			reqRecorder.Body.String(), expected)
	}

}

func TestGetCustomerById(t *testing.T) {

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	rp := &mock.CustomerModel{}

	cs := service.DefaultCustomerService{
		Repository: rp,
		Logger:     logger,
		Helpers:    helpers.Helpers{},
		Errors:     helpers.Errors{},
	}
	app := &application{
		logger:          logger,
		customerService: cs,
	}

	req, err := http.NewRequest("GET", "/customers/1", nil)
	if err != nil {
		t.Fatal(err)
	}
	reqRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(app.customerService.GetCustomers)
	handler.ServeHTTP(reqRecorder, req)

	// Test whether the response is OK
	if status := reqRecorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Test the expected response body
	expected := `{"customers":[{"ID":1,"Name":"Bob","Age":22,"Country":"MT",`
	if !strings.Contains(reqRecorder.Body.String(), expected) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			reqRecorder.Body.String(), expected)
	}

}
