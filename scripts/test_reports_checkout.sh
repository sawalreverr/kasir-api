#!/bin/bash

BASE_URL="http://localhost:8080"

echo "1. GET /report/today (get today's sales report)"
curl -i "${BASE_URL}/report/today"
echo -e "\n"

echo "2. GET /report?start_date=2026-02-06&end_date=2026-02-09 (get date range report)"
curl -i "${BASE_URL}/report?start_date=2026-01-01&end_date=2026-02-01"
echo -e "\n"

echo "3. GET /report with invalid date format (error handling test)"
curl -i "${BASE_URL}/report?start_date=invalid&end_date=2026-02-01"
echo -e "\n"

echo "4. GET /report with start_date after end_date (error handling test)"
curl -i "${BASE_URL}/report?start_date=2026-03-01&end_date=2026-02-01"
echo -e "\n"

echo "5. POST /checkout (create transaction with 2 items)"
curl -i -X POST \
  -H "Content-Type: application/json" \
  -d '{
    "items": [
      {"product_id": 1, "quantity": 2},
      {"product_id": 2, "quantity": 1}
    ]
  }' \
  "${BASE_URL}/checkout"
echo -e "\n"

echo "6. POST /checkout with empty items (error handling test)"
curl -i -X POST \
  -H "Content-Type: application/json" \
  -d '{
    "items": []
  }' \
  "${BASE_URL}/checkout"
echo -e "\n"

echo "7. POST /checkout with invalid product_id (error handling test)"
curl -i -X POST \
  -H "Content-Type: application/json" \
  -d '{
    "items": [
      {"product_id": 999, "quantity": 1}
    ]
  }' \
  "${BASE_URL}/checkout"
echo -e "\n"

echo "8. POST /checkout with zero quantity (error handling test)"
curl -i -X POST \
  -H "Content-Type: application/json" \
  -d '{
    "items": [
      {"product_id": 1, "quantity": 0}
    ]
  }' \
  "${BASE_URL}/checkout"
echo -e "\n"

echo "9. GET /report/today again (verify transaction was created)"
curl -i "${BASE_URL}/report/today"
echo -e "\n"
