.PHONY: build test clean run install

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
