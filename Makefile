PROJECT_NAME = oapicodegen-example

# Golang
compile :
	CGO_ENABLED=0 go build -o ./bin/${PROJECT_NAME} ./cmd/app

go:	compile
	./bin/${PROJECT_NAME}

lint:
	golangci-lint run ./... --fix

fmt:
	gofmt -s -w .
	goimports -l -w .

test:
	go test -v ./internal/...

# Docker
down:
	docker compose down

build:
	docker compose build

up:
	docker compose up

run: down build up

stop:
	docker stop $$(docker ps -a -q)

remove:
	docker rm $$(docker ps -a -q)

clean: stop remove
	docker system prune

openapi:
	cd ./api/v1/ && bash create_schema.sh

gen: openapi
	go generate ./...
	go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen -o internal/gen/card/server_types.go -generate fiber-server,types -package cardgen api/v1/card/schema.yaml
	go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen -o internal/gen/bank/server_types.go -generate fiber-server,types -package bankgen api/v1/bank/schema.yaml
