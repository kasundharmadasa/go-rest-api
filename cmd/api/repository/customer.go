package repository

import (
	"sample.api.kasun.com/pkg/models"
)

type CustomerRepository interface {
	Insert(name string, age int, country string, items []string, status string) (int, error)
	GetCustomerById(id int64) (*models.Customer, error)
	GetCustomers() ([]models.Customer, error)
	UpdateCustomerStatus(id int, status string) (int, error)
}
