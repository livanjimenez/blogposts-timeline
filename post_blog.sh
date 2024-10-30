#!/bin/bash

# Define the JSON payload
JSON_PAYLOAD='{
  "title": "My First Blog",
  "content": "This is the content of my first blog post."
}'

# Make the POST request using curl
curl -X POST http://localhost:8080/postblog \
-H "Content-Type: application/json" \
-d "$JSON_PAYLOAD"
