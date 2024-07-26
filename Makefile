.PHONY: run tidy lint build

run:
	go run main.go

tidy:
	go mod tidy

lint:
	golangci-lint run

hello:
	echo "Hello All"

build:
	go build -o main .