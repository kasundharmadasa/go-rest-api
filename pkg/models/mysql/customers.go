package mysql

import (
	"database/sql"

	"sample.api.kasun.com/pkg/models"
)

// Define a CustomerModel type which wraps a sql.DB connection pool.
type CustomerModel struct {
	DB *sql.DB
}

func (m *CustomerModel) Insert(name string, age int, country string, items []string) (int, error) {

	customerInsertStmt := `INSERT INTO CUSTOMERS (NAME, AGE, COUNTRY) VALUES(?, ?, ?)`

	result, err := m.DB.Exec(customerInsertStmt, name, age, country)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	itemsInsertStmt := `INSERT INTO CUSTOMER_ITEMS (CUSTOMER_ID, ITEM) VALUES(?, ?)`

	for _, item := range items {
		_, err := m.DB.Exec(itemsInsertStmt, id, item)
		if err != nil {
			return 0, err
		}
	}

	return 0, nil
}

func (m *CustomerModel) GetCustomerById(id int64) (*models.Customer, error) {

	stmt := `SELECT ID, NAME, AGE, COUNTRY, ITEM FROM CUSTOMERS INNTER JOIN CUSTOMER_ITEMS
	 on ID = CUSTOMER_ITEMS.CUSTOMER_ID where ID=?`

	rows, err := m.DB.Query(stmt, id)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	customer := &models.Customer{}
	for rows.Next() {

		var customerEntry struct {
			ID      int64
			Name    string
			Age     int32
			Country string
			Item    string
		}

		err = rows.Scan(&customerEntry.ID, &customerEntry.Name, &customerEntry.Age,
			&customerEntry.Country, &customerEntry.Item)

		if err != nil {
			return nil, err
		}

		if customer.ID == customerEntry.ID {
			customer.Items = append(customer.Items, customerEntry.Item)
		} else {
			customer.ID = customerEntry.ID
			customer.Name = customerEntry.Name
			customer.Age = customerEntry.Age
			customer.Country = customerEntry.Country
			customer.Items = append(customer.Items, customerEntry.Item)
		}
	}
	return customer, nil
}

func (m *CustomerModel) GetCustomers() ([]models.Customer, error) {

	stmt := `SELECT ID, NAME, AGE, COUNTRY, ITEM FROM CUSTOMERS INNTER JOIN CUSTOMER_ITEMS
	 on ID = CUSTOMER_ITEMS.CUSTOMER_ID`

	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	customers := []models.Customer{}
	for rows.Next() {

		customer := &models.Customer{}

		var customerEntry struct {
			ID      int64
			Name    string
			Age     int32
			Country string
			Item    string
		}

		err = rows.Scan(&customerEntry.ID, &customerEntry.Name, &customerEntry.Age,
			&customerEntry.Country, &customerEntry.Item)

		if err != nil {
			return nil, err
		}

		var found bool = false
		for index, element := range customers {
			if element.ID == customerEntry.ID {
				customers[index].Items = append(customers[index].Items, customerEntry.Item)
				found = true
			}
		}
		if !found {
			customer := &models.Customer{
				ID:      customerEntry.ID,
				Name:    customerEntry.Name,
				Age:     customerEntry.Age,
				Country: customerEntry.Country,
				Items:   append(customer.Items, customerEntry.Item),
			}

			customers = append(customers, *customer)
		}
	}
	return customers, nil
}
