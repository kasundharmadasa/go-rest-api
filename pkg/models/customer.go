package models

// Customer defines a struct to hold the customer details
type Customer struct {
	ID      int64
	Name    string
	Age     int32
	Country string
	Items   []string
	Status  string
}
