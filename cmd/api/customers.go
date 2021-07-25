package main

import (
	"net/http"

	"sample.api.kasun.com/pkg/models"
)

var cust = map[int64]*models.Customer{}

// Create a customer
func (app *application) createCustomerHandler(w http.ResponseWriter, r *http.Request) {

	// Creating a separate struct to hold the request as the customer struct has additional ID params
	var input struct {
		Name    string   `json:"name"`
		Age     int32    `json:"age"`
		Country string   `json:"country"`
		Items   []string `json:"items"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	customer := &models.Customer{
		Name:    input.Name,
		Age:     input.Age,
		Country: input.Country,
		Items:   input.Items,
	}

	_, dberr := app.customers.Insert(customer.Name, int(customer.Age), customer.Country, customer.Items)

	if dberr != nil {
		app.logger.Println(dberr)
	}

}

// Get customer details for a given ID
func (app *application) getCustomerByIDHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	customers, dberr := app.customers.GetCustomerById(id)

	if dberr != nil {
		app.logger.Println(dberr)
	}

	if customers == nil {
		http.NotFound(w, r)
		return
	}

	err = app.writeJSON(w, http.StatusOK, customers, nil)
	if err != nil {
		app.logger.Println(err)
		http.Error(w, "The server encountered a problem and could not process your request",
			http.StatusInternalServerError)
	}
}

// Get the list of all the customers
func (app *application) getCustomersHandler(w http.ResponseWriter, r *http.Request) {

	customers := []models.Customer{}
	for _, customer := range cust {
		customers = append(customers, *customer)
	}

	customers, dberr := app.customers.GetCustomers()

	if dberr != nil {
		app.logger.Println(dberr)
	}

	err := app.writeJSON(w, http.StatusOK, customers, nil)
	if err != nil {
		app.logger.Println(err)
		http.Error(w, "The server encountered a problem and could not process your request",
			http.StatusInternalServerError)
	}
}
