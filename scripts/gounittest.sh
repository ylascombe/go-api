#!/usr/bin/env bash

if [ -z $1 ];
then
  echo "==> Running all unit test suite"
  for pkg in $(go list ./...);
  do
    go test "$pkg"
    echo "------------------------"
  done
else
  echo "==> Running $1 unit test suite"
  go test "$1"
fi
