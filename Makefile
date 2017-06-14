TEST_FILES=$(go list ./...)
BIN    = $(GOPATH)/bin
GOLINT = $(BIN)/golint

$(BIN)/golint: | $(BASE)

default: vet lintall testall coverstat build

build:
	@go get
	@go build

# test runs the unit tests
# make testall
testall:
	@sh -c "'$(CURDIR)/scripts/gounittest.sh'"

# test run component
# make test TEST=arc-api
test:
	@sh -c "'$(CURDIR)/scripts/gounittest.sh' $(TEST)"

cover:
	@go tool cover 2>/dev/null; if [ $$? -eq 3 ]; then \
		go get -u golang.org/x/tools/cmd/cover; \
	fi
	go test $(TEST) -coverprofile=coverage.out
	go tool cover -html=coverage.out
	rm coverage.out

coverstat:
	@sh -c "'$(CURDIR)/scripts/gocover.sh'"

vet:
	@echo '[ CMD ]=> go vet $$(go list ./... | grep -v /terraform/vendor/)'
	@go vet $$(go list ./... | grep -v /terraform/vendor/) ; if [ $$? -eq 1 ]; then \
		echo ""; \
		echo "Vet found suspicious constructs. Please check the reported constructs"; \
		echo "and fix them if necessary before submitting the code for review."; \
		exit 1; \
	fi

# make lint PKG=arc-api
lint:
	@golint 2>/dev/null; if [ $$? -eq 3 ]; then \
	  go get -u github.com/golang/lint/golint; \
	fi
	golint -set_exit_status $(PKG)

lintall:
	@sh -c "'$(CURDIR)/scripts/golint.sh'"

clean:
	rm -f arc-api

dependencies:
	go get

run:
	go run main.go

tests:
	go test ./...

debug_tools:
	go get github.com/derekparker/delve/cmd/dlv