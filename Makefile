.SILENT:

.PHONY: fmt lint test race run dc_up dc_create_network dc_down

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

dc_down:
	docker compose down

goose_up:
	goose -dir ./db/migrations postgres \
	"postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DBNAME}?sslmode=${POSTGRES_SSLMODE}" up

goose_down:
	goose -dir ./db/migrations postgres \
	"postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DBNAME}?sslmode=${POSTGRES_SSLMODE}" down

goose_status:
	goose -dir ./db/migrations postgres \
	"postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DBNAME}?sslmode=${POSTGRES_SSLMODE}" status

proto_install:
	apt update && apt install -y protobuf-compiler \
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest \
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest \
	protoc --version

proto_gen_sso:
	protoc -I ./protos/proto --go_out=./protos/gen --go_opt=paths=source_relative \
	--go-grpc_out=./protos/gen --go-grpc_opt=paths=source_relative \
	--grpc-gateway_out=./protos/gen --grpc-gateway_opt=paths=source_relative \
	./protos/proto/sso_service/sso_auth.proto

# default 
.DEFAULT_GOAL := run