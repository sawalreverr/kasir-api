#!/bin/bash

BASE_URL="http://localhost:8080"
DUMMY_ID=3

echo "1. GET /products (get all products)"
curl -i "${BASE_URL}/products"
echo -e "\n"

echo "2. POST /products (create dummy product)"
curl -i -X POST \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Keyboard",
    "price": 300000,
    "stock": 15,
    "category_ids": ["1"]
  }' \
  "${BASE_URL}/products"
echo -e "\n"

echo "3. GET /products/${DUMMY_ID} (get dummy product)"
curl -i "${BASE_URL}/products/${DUMMY_ID}"
echo -e "\n"

echo "4. PUT /products/${DUMMY_ID} (update dummy product)"
curl -i -X PUT \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Mechanical Keyboard",
    "price": 450000,
    "stock": 10,
    "category_ids": ["1"]
  }' \
  "${BASE_URL}/products/${DUMMY_ID}"
echo -e "\n"

echo "5. DELETE /products/${DUMMY_ID} (delete dummy product)"
curl -i -X DELETE "${BASE_URL}/products/${DUMMY_ID}"
echo -e "\n"

echo "6. GET /products (check main data still intact)"
curl -i "${BASE_URL}/products"
echo -e "\n"
