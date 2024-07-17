.PHONY: run tidy lint

run:
	go run main.go

tidy:
	go mod tidy

lint:
	golangci-lint run

hello:
	echo "Hello All"