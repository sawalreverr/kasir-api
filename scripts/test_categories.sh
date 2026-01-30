#!/bin/bash

BASE_URL="http://localhost:8080"
DUMMY_ID=4

echo "1. GET /categories (get all categories)"
curl -i "${BASE_URL}/categories"
echo -e "\n"

echo "2. POST /categories (create dummy category)"
curl -i -X POST \
  -H "Content-Type: application/json" \
  -d '{"name":"Accessories","description":"Accessories category"}' \
  "${BASE_URL}/categories"
echo -e "\n"

echo "3. GET /categories/${DUMMY_ID} (get dummy category)"
curl -i "${BASE_URL}/categories/${DUMMY_ID}"
echo -e "\n"

echo "4. PUT /categories/${DUMMY_ID} (update dummy category)"
curl -i -X PUT \
  -H "Content-Type: application/json" \
  -d '{"name":"Accessories Updated","description":"Updated accessories"}' \
  "${BASE_URL}/categories/${DUMMY_ID}"
echo -e "\n"

echo "5. DELETE /categories/${DUMMY_ID} (delete dummy category)"
curl -i -X DELETE "${BASE_URL}/categories/${DUMMY_ID}"
echo -e "\n"

echo "6. GET /categories (check main data)"
curl -i "${BASE_URL}/categories"
echo -e "\n"
