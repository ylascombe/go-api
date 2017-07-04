#!/bin/bash

curl http://localhost:8090/v1/environments/ | jq
curl http://localhost:8090/v1/environments/LOCAL -X POST
idEnv1=$(curl http://localhost:8090/v1/environments/ | jq .[0].id)

curl http://localhost:8090/v1/users/ | jq
curl http://localhost:8090/v1/users/ -X POST -H "Content-Type: application/json" -d @user1.json
idUser1=$(curl http://localhost:8090/v1/users/ | jq .[0].ID)

curl http://localhost:8090/v1/users/ -X POST -H "Content-Type: application/json" -d @user2.json
idUser2=$(curl http://localhost:8090/v1/users/ | jq .[1].ID)

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

curl http://localhost:8090/v1/environments/LOCAL/access/ | jq
curl http://localhost:8090/v1/environments/LOCAL/access/$idUser1 -X POST -H "Content-Type: application/json" -d @tmp-access1.json
curl http://localhost:8090/v1/environments/LOCAL/access/$idUser2 -X POST -H "Content-Type: application/json" -d @tmp-access2.json


curl http://127.0.0.1:8090/v1/ssh-keys/LOCAL

curl http://127.0.0.1:8090/v1/teams/ | jq
curl http://127.0.0.1:8090/v1/teams/ -X POST -H "Content-Type: application/json" -d @ft1.json
idFT1=$(curl http://localhost:8090/v1/teams/ | jq .[0].ID)

curl http://127.0.0.1:8090/v1/teams/colis360/user/ | jq
curl http://127.0.0.1:8090/v1/teams/colis360/user/$idUser1 -X POST -H "Content-Type: application/json" -d @tmp-membership1.json
