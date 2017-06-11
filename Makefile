
BIN    = $(GOPATH)/bin
GOLINT = $(BIN)/golint

$(BIN)/golint: | $(BASE)
	go get github.com/golang/lint/golint.

dependencies:
	go get github.com/lib/pq
	#go get github.com/go-xorm/xorm
	go get -u github.com/jinzhu/gorm

run:
	go run main.go
