all: lint

lint:
	env GOGC=25 golangci-lint run --fix -j 8 -v ./...