#!/bin/bash

BASE_URL="http://localhost:8080"

echo "1. GET /categories (get all categories)"
curl -i "${BASE_URL}/categories"
echo -e "\n"

echo "2. POST /categories (create new category)"
curl -i -X POST -H "Content-Type: application/json" -d '{"name":"clothing", "description":"Clothings"}' "${BASE_URL}/categories"
echo -e "\n"

echo "3. GET /categories/1 (get category by id 1)"
curl -i "${BASE_URL}/categories/1"
echo -e "\n"

echo "4. PUT /categories/1 (update category id 1)"
curl -i -X PUT -H "Content-Type: application/json" -d '{"id": 1, "name":"headset", "description":"Headset"}' "${BASE_URL}/categories/1"
echo -e "\n"

echo "5. DELETE /categories/2 (delete category id 2)"
curl -i -X DELETE "${BASE_URL}/categories/2"
echo -e "\n"

echo "6. GET /categories (check changes)"
curl -i "${BASE_URL}/categories"
echo -e "\n"
