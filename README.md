# Basic GO API

## Task Session 1
Implementation Category CRUD [DONE]

## Task Session 2
1. Pindah categories temen-temen ke layered architecture [DONE]
2. Challange (Optional): Explore Join, tambah category_id ke table products, setiap product mempunyai kategory, dan ketika Get Detail return category.name dari product. [DONE]

## Task Session 3
1. Tambah `/checkout` untuk transaction product [DONE]
2. Sales summary hari ini dan sales date range (optional) [DONE]

## How to Run

### Run the server
```bash
go run cmd/server/main.go
```

### Run the tests
```bash
./scripts/test_categories.sh
./scripts/test_products.sh
./scripts/test_reports_checkout.sh
```

### Endpoint API
`Categories`
* **GET** `/categories` -> Get all categories
* **POST** `/categories` -> Create category
* **PUT** `/categories/{id}` -> Update category
* **GET** `/categories/{id}` -> Get category by id
* **DELETE** `/categories/{id}` -> Delete category

`Products`
* **GET** `/products` -> Get all products
* **POST** `/products` -> Create product
* **PUT** `/products/{id}` -> Update product
* **GET** `/products/{id}` -> Get product by id
* **DELETE** `/products/{id}` -> Delete product

`Transaction`
* **POST** `/checkout` -> Create transaction (checkout)

`Report`
* **GET** `/report/today` -> Get today's sales report
* **GET** `/report` -> Get sales report for date range (requires `start_date` and `end_date` query params)

### Example Usage
```bash
$ curl -i -X GET "http://localhost:8080/categories"
HTTP/1.1 200 OK
Content-Type: application/json
Date: Thu, 22 Jan 2026 15:37:27 GMT
Content-Length: 201

{"success":true,"message":"success","data":[{"id":"1","name":"Electronics","description":"Electronic devices"},{"id":"2","name":"Books","description":"All kinds of books"},{"id":"3","name":"Clothing","description":"Wearable items"}]}
```
