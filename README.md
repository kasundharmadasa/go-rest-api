
# REST API example application

This is a sample REST API implementation to manage customer details.

# REST API

The REST API to the example app is described below.


## Create a new Customer

### Request

`POST /customers`

    curl -i -d '{"name":"Bill","age":30,"country":"AU", "items": ["Keyboard","Speaker"]}' localhost:4000/customers

### Response

	HTTP/1.1 201 Created
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

	{"customers":[{"ID":1,"Name":"Moana","Age":26,"Country":"BG","Items":["Laptop","Speaker"]},{"ID":2,"Name":"Bill","Age":30,"Country":"AU","Items":["Keyboard","Speaker"]}]}
	

## Get a specific Customer

### Request

`GET /customers/id`

     curl -i localhost:4000/customers/1 

### Response

	HTTP/1.1 200 OK
	Content-Type: application/json
	Date: Wed, 21 Jul 2021 15:59:43 GMT
	Content-Length: 91

	{"customers":{"ID":1,"Name":"Moana","Age":26,"Country":"BG","Items":["Laptop","Speaker"]}}
