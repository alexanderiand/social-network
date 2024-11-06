.SILENT:

.PHONY: fmt lint test race run dc_up

include .env 
export 

# fmt
fmt: 
	go fmt ./...
# lint
lint: fmt 
	go vet ./...

# test
test: lint
	go  test -v -cover ./...

# race
race: test
	go test -v -race ./...

# run
run: race 
	go run cmd/social-network/main.go 

dc_create_network:
	docker network create social_network_net

dc_up:
	docker compose up -d

# default 
.DEFAULT_GOAL := dc_up