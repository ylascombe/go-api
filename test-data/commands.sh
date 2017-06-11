#!/bin/bash

curl http://localhost:8080/v1/user | jq
curl http://localhost:8080/v1/user -X POST -H "Content-Type: application/json" -d @user1.json
