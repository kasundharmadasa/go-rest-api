package main

import (
	"bytes"
	"log"
	"os"

	"net/http"
	"net/http/httptest"
	"testing"

	"sample.api.kasun.com/pkg/models/mock"
)

func TestGetCustomers(t *testing.T) {

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)
	app := &application{
		logger:    logger,
		customers: &mock.CustomerModel{},
	}

	req, err := http.NewRequest("GET", "/customers", nil)
	if err != nil {
		t.Fatal(err)
	}
	reqRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(app.getCustomersHandler)
	handler.ServeHTTP(reqRecorder, req)

	// Test whether the response is OK
	if status := reqRecorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Test the expected response body
	expected := []byte(`{"customers":[{"ID":1,"Name":"Bob","Age":22,"Country":"MT",` +
		`"Items":["Mouse","Keyboard"]}]}`)
	if !bytes.Contains(reqRecorder.Body.Bytes(), expected) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			reqRecorder.Body.String(), expected)
	}

}

func TestGetCustomerById(t *testing.T) {

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)
	app := &application{
		logger:    logger,
		customers: &mock.CustomerModel{},
	}

	req, err := http.NewRequest("GET", "/customers/1", nil)
	if err != nil {
		t.Fatal(err)
	}
	reqRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(app.getCustomersHandler)
	handler.ServeHTTP(reqRecorder, req)

	// Test whether the response is OK
	if status := reqRecorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Test the expected response body
	expected := []byte(`{"customers":[{"ID":1,"Name":"Bob","Age":22,"Country":"MT",` +
		`"Items":["Mouse","Keyboard"]}]}`)
	if !bytes.Contains(reqRecorder.Body.Bytes(), expected) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			reqRecorder.Body.String(), expected)
	}

}

func TestCreateCustomer(t *testing.T) {

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)
	app := &application{
		logger:    logger,
		customers: &mock.CustomerModel{},
	}

	req, err := http.NewRequest("POST", "/customers", nil)
	if err != nil {
		t.Fatal(err)
	}
	reqRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(app.getCustomersHandler)
	handler.ServeHTTP(reqRecorder, req)

	// Test whether the response is OK
	if status := reqRecorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}
