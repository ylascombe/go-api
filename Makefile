
BIN    = $(GOPATH)/bin
GOLINT = $(BIN)/golint

$(BIN)/golint: | $(BASE)

dependencies:
	go get github.com/lib/pq
	#go get github.com/go-xorm/xorm
	go get -u github.com/jinzhu/gorm

run:
	go run main.go

tests:
	go test ./...
