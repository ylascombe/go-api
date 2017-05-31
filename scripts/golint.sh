#!/usr/bin/env bash

golint 2>/dev/null; if [ $? -eq 3 ]; then
  go get -u github.com/golang/lint/golint;
fi

echo "==> Running linter"
for pkg in $(go list ./...);
do
  golint -set_exit_status "$pkg"
  echo "------------------------"
done
