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

curl http://127.0.0.1:8080/v1/featureTeam | jq
curl http://127.0.0.1:8080/v1/featureTeam -X POST -H "Content-Type: application/json" -d @ft1.json
idFT1=$(curl http://localhost:8080/v1/featureTeam | jq .featureTeams[0].ID)


cat <<EOF > tmp-membership1.json
{
  "ID": 27,
  "ApiUserID": $idUser1,
  "FeatureTeamID": $idFT1
}
EOF
curl http://127.0.0.1:8080/v1/membership/colis360 | jq
curl http://127.0.0.1:8080/v1/membership/colis360 -X POST -H "Content-Type: application/json" -d @tmp-membership1.json
