
# REST API example application

This is a sample REST API implementation to manage customer details.

# REST API

The REST API to the example app is described below.


# Setup

## Setting up database
 1. Create a MYSQL database and execute `dbscripts/mysql.sql`
 2. Create a .envrc file in $PROJECT_HOME directory and configure 'CUSTOMER_API_DB_DSN' with database details 
	
		export CUSTOMER_API_DB_DSN=root:root@tcp(127.0.0.1:3306)/go_customer

## Setting up streams in Apache Kafka

1. Setup Apache Kafka cluster along with ksqlDB
2. Create `user_transactions` and `possible_user_anomalies` streams as described in `streams/kafka.sql`
3. Send transaction events as descibed in `streams/kafka.sql` 

## Create a new Customer

### Request

`POST /customers`

    curl -i -d '{"name":"Bill","age":30,"country":"MT", "items": ["Keyboard","Speaker"]}' localhost:4000/customers

### Response

	HTTP/1.1 200 OK
	Date: Wed, 21 Jul 2021 15:57:01 GMT
	Content-Length: 0	


## Get list of Customers

### Request

`GET /customers`

    curl -i localhost:4000/customers

### Response

	HTTP/1.1 200 OK
	Content-Type: application/json
	Date: Wed, 21 Jul 2021 15:58:38 GMT
	Content-Length: 171

	{"customers":[{"ID":1,"Name":"Moana","Age":26,"Country":"BG","Items":["Laptop","Speaker"]},{"ID":2,"Name":"Bill","Age":30,"Country":"MT","Items":["Keyboard","Speaker"]}]}
	

## Get a specific Customer

### Request

`GET /customers/id`

     curl -i localhost:4000/customers/1 

### Response

	HTTP/1.1 200 OK
	Content-Type: application/json
	Date: Wed, 21 Jul 2021 15:59:43 GMT
	Content-Length: 91

	{"customers":{"ID":1,"Name":"Moana","Age":26,"Country":"MT","Items":["Laptop","Speaker"]}}
