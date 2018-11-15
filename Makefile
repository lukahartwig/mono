BINARY=mono

build:
	go build -o ${BINARY}

dependencies:
	go mod download

test:
	go test -v ./...

testCI:
	go test -v -race -coverprofile=coverage.txt ./...

.PHONY: build test testCI
