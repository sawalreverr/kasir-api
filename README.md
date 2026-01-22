# Basic GO API

## Task Session 1
Implementation Category CRUD

### Model
**Category**
* ID
* Name
* Description

### Endpoint yang Wajib Ada
* **GET** `/categories` -> Get all categories
* **POST** `/categories` -> Create category
* **PUT** `/categories/{id}` -> Update category 
* **GET** `/categories/{id}` -> Get category by id
* **DELETE** `/categories/{id}` -> Delete category

## How to Run

### Run the server
```bash
go run main.go
```

### Run the tests
```bash
./test_api.sh
```

### Example Usage
```bash
$ curl -i -X GET "http://localhost:8080/categories"
HTTP/1.1 200 OK
Content-Type: application/json
Date: Thu, 22 Jan 2026 15:37:27 GMT
Content-Length: 201

{"success":true,"message":"all categories","data":[{"id":1,"name":"food","description":"Foods"},{"id":2,"name":"electronics","description":"Electronics"},{"id":3,"name":"book","description":"Books"}]}
```
