package service

import (
	"net/http"
)

type CustomerService interface {
	CreateCustomer(w http.ResponseWriter, r *http.Request)
	GetCustomerByID(w http.ResponseWriter, r *http.Request)
	GetCustomers(w http.ResponseWriter, r *http.Request)
	BlockCustomer(id int)
}
