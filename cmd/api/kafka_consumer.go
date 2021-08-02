package main

import (
	"encoding/json"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"sample.api.kasun.com/pkg/models"
)

func (app *application) consumeKafka() {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost",
		"group.id":          "possible_user_anomalies",
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		app.logger.Println(err)
	}

	c.SubscribeTopics([]string{"possible_user_anomalies"}, nil)

	for {
		msg, err := c.ReadMessage(-1)
		if err == nil {
			//fmt.Printf("Message on %s: \n", string(msg.Value))
			var message models.KafkaCustomerMessage
			json.Unmarshal([]byte(string(msg.Value)), &message)

			// update customer status to BLOCKED in database as possible anomaly detected
			app.customerService.BlockCustomer(message.UserID)
		} else {
			// The client will automatically try to recover from all errors.
			app.logger.Printf("Consumer error: %v (%v)\n", err, msg)
			break
		}

	}

	c.Close()
}
