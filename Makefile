pkgs          = $(shell go list ./... | grep -v /tests | grep -v /vendor/ | grep -v /common/)

test:
	@echo " >> running tests"
	@go test  -cover $(pkgs)

race:
	@echo " >> running tests with race"
	@go test  -cover -race $(pkgs)

run:
	ENV=development go run main.go serve-http

hot:
	gin -p 9001 -a 8080 serve-http

install:
	@go mod download

.PHONY: test clean
