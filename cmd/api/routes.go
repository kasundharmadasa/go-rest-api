package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() *httprouter.Router {

	router := httprouter.New()

	router.HandlerFunc(http.MethodPost, "/customers", app.createCustomerHandler)
	router.HandlerFunc(http.MethodGet, "/customers", app.getCustomersHandler)
	router.HandlerFunc(http.MethodGet, "/customers/:id", app.getCustomerByIDHandler)

	// Return the httprouter instance.
	return router
}
