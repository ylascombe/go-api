
BIN    = $(GOPATH)/bin
GOLINT = $(BIN)/golint

$(BIN)/golint: | $(BASE)

dependencies:
	go get

run:
	go run main.go

tests:
	go test ./...

debug_tools:
	go get github.com/derekparker/delve/cmd/dlv    
