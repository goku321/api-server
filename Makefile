PROJECTNAME := $(shell basename "$(PWD)")
all : clean fmt test build run

build:
	@echo " > Building api server..."
	@go build $(LDFLAGS) -o $(PROJECTNAME)

start-server:
	docker-compose up --detach

stop-server:
	docker-compose down

test:
	go test -count=1 ./... -v

test-e2e:
	@docker-compose up -d postgres
	go test -build=integration -v
	docker-compose down

clean:
	go clean

fmt:
	go fmt ./...