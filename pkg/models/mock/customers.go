package mock

import (
	"sample.api.kasun.com/pkg/models"
)

type CustomerModel struct {
}

func (m *CustomerModel) Insert(name string, age int, country string, items []string) (int, error) {

	return 0, nil
}

func (m *CustomerModel) GetCustomerById(id int64) (*models.Customer, error) {

	customer := &models.Customer{
		ID:      1,
		Name:    "Bob",
		Age:     22,
		Country: "MT",
		Items:   []string{"Mouse", "Keyboard"},
	}

	return customer, nil
}

func (m *CustomerModel) GetCustomers() ([]models.Customer, error) {

	customers := []models.Customer{}

	customer := &models.Customer{
		ID:      1,
		Name:    "Bob",
		Age:     22,
		Country: "MT",
		Items:   []string{"Mouse", "Keyboard"},
	}

	customers = append(customers, *customer)

	return customers, nil
}
