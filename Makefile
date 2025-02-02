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

gen: 
	go generate ./...
	go generate -v tools.go

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

