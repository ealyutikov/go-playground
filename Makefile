build:
	@go build -o bin/goblockchain

run: build
	@./bin/docker

test:
	@go test -v ./...

