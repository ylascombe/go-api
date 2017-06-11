#!/bin/bash

curl http://localhost:8080/v1/environment | jq
curl http://localhost:8080/v1/environment/UNKNOWN -X POST

curl http://localhost:8080/v1/user | jq
curl http://localhost:8080/v1/user -X POST -H "Content-Type: application/json" -d @user1.json

curl http://localhost:8080/v1/environmentAccess | jq
curl http://localhost:8080/v1/environmentAccess -X POST -H "Content-Type: application/json" -d @access1.json

curl http://localhost:8080/v1/environmentAccess/SPECIFIC | jq
curl http://localhost:8080/v1/environmentAccess/SPECIFIC/user/99 -X PUT -H "Content-Type: application/json"
