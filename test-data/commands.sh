#!/bin/bash

curl http://localhost:8080/v1/environment | jq
curl http://localhost:8080/v1/environment/LOCAL -X POST
idEnv1=$(curl http://localhost:8080/v1/environment | jq .[0].ID)

curl http://localhost:8080/v1/user | jq
curl http://localhost:8080/v1/user -X POST -H "Content-Type: application/json" -d @user1.json
idUser1=$(curl http://localhost:8080/v1/user | jq .users[0].ID)

curl http://localhost:8080/v1/user -X POST -H "Content-Type: application/json" -d @user2.json
idUser2=$(curl http://localhost:8080/v1/user | jq .users[1].ID)

cat <<EOF > tmp-access1.json
{
	"ApiUserID": $idUser1,
	"EnvironmentID": $idEnv1
}
EOF

cat <<EOF > tmp-access2.json
{
	"ApiUserID": $idUser2,
	"EnvironmentID": $idEnv1
}
EOF

curl http://localhost:8080/v1/environmentAccess/LOCAL | jq
curl http://localhost:8080/v1/environmentAccess/LOCAL/user/$idUser1 -X PUT -H "Content-Type: application/json" -d @tmp-access1.json
curl http://localhost:8080/v1/environmentAccess/LOCAL/user/$idUser2 -X PUT -H "Content-Type: application/json" -d @tmp-access2.json


curl http://127.0.0.1:8080/v1/sshKeys/LOCAL