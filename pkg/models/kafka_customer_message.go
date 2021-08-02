package models

// struct holding the customer messages received from kafka
type KafkaCustomerMessage struct {
	UserID   int     `json:"user_id"`
	Attempts int     `json:"n_attempts"`
	Amount   float64 `json:"total_amount"`
}
