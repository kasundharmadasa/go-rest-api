package service

import (
	"log"
	"net/http"

	"sample.api.kasun.com/cmd/api/helpers"
	"sample.api.kasun.com/cmd/api/repository"
	"sample.api.kasun.com/pkg/models"
)

type DefaultCustomerService struct {
	Repository repository.CustomerRepository
	Logger     *log.Logger
	Helpers    helpers.Helpers
	Errors     helpers.Errors
}

var cust = map[int64]*models.Customer{}

// Create a customer
func (customerService DefaultCustomerService) CreateCustomer(w http.ResponseWriter, r *http.Request) {

	// Creating a separate struct to hold the request as the customer struct has additional ID params
	var input struct {
		Name    string   `json:"name"`
		Age     int32    `json:"age"`
		Country string   `json:"country"`
		Items   []string `json:"items"`
	}

	err := customerService.Helpers.ReadJSON(w, r, &input)
	if err != nil {
		customerService.Errors.BadRequestResponse(w, r, err)
		return
	}

	customer := &models.Customer{
		Name:    input.Name,
		Age:     input.Age,
		Country: input.Country,
		Items:   input.Items,
		Status:  "ACTIVE",
	}

	_, dberr := customerService.Repository.Insert(customer.Name, int(customer.Age), customer.Country, customer.Items, customer.Status)

	if dberr != nil {
		customerService.Logger.Println(dberr)
	}

}

// Get customer details for a given ID
func (customerService DefaultCustomerService) GetCustomerByID(w http.ResponseWriter, r *http.Request) {
	id, err := customerService.Helpers.ReadIDParam(r)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	customers, dberr := customerService.Repository.GetCustomerById(id)

	if dberr != nil {
		customerService.Logger.Println(dberr)
	}

	if customers == nil {
		http.NotFound(w, r)
		return
	}

	err = customerService.Helpers.WriteJSON(w, http.StatusOK, customers, nil)
	if err != nil {
		customerService.Logger.Println(err)
		http.Error(w, "The server encountered a problem and could not process your request",
			http.StatusInternalServerError)
	}
}

// Get the list of all the customers
func (customerService DefaultCustomerService) GetCustomers(w http.ResponseWriter, r *http.Request) {

	customers := []models.Customer{}
	for _, customer := range cust {
		customers = append(customers, *customer)
	}

	customers, dberr := customerService.Repository.GetCustomers()

	if dberr != nil {
		customerService.Logger.Println(dberr)
	}

	err := customerService.Helpers.WriteJSON(w, http.StatusOK, customers, nil)
	if err != nil {
		customerService.Logger.Println(err)
		http.Error(w, "The server encountered a problem and could not process your request",
			http.StatusInternalServerError)
	}
}

func (customerService DefaultCustomerService) BlockCustomer(id int) {

	_, err := customerService.Repository.UpdateCustomerStatus(id, "BLOCKED")
	if err != nil {
		customerService.Logger.Println(err)
	}
}
