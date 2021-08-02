package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() *httprouter.Router {

	router := httprouter.New()

	router.HandlerFunc(http.MethodPost, "/customers", app.customerService.CreateCustomer)
	router.HandlerFunc(http.MethodGet, "/customers", app.customerService.GetCustomers)
	router.HandlerFunc(http.MethodGet, "/customers/:id", app.customerService.GetCustomerByID)

	// Return the httprouter instance.
	return router
}
