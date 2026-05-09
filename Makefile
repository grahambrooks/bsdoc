.PHONY: build test clean run install update-all update fmt check

build:
	go build -o bin/bsdoc ./cmd/bsdoc

run:
	go run ./cmd/bsdoc

test:
	go test ./...

clean:
	rm -rf bin/

install:
	go install ./cmd/bsdoc

update-all: update fmt check test

update:
	go get -u ./...
	go mod tidy

fmt:
	gofmt -w .

check:
	go vet ./...
