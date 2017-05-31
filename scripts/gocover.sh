#!/usr/bin/env bash

go tool cover 2>/dev/null; if [ $? -eq 3 ]; then
  go get -u golang.org/x/tools/cmd/cover;
fi

echo "==> Running all coverage unit test suite"
for pkg in $(go list ./...);
do
  go test -coverprofile "$pkg"
  echo "------------------------"
done
