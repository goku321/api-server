PROJECTNAME := $(shell basename "$(PWD)")
all : clean fmt test build run

build:
	@echo " > Building api server..."
	@go build $(LDFLAGS) -o $(PROJECTNAME)

# start-server:
# 	docker run --network mynetwork --name postgres -e POSTGRES_PASSWORD=password -d -p 6432:5432 postgres
# 	docker run --network mynetwork --name api-server -e DB_CONN_STR=postgres://postgres:password@postgres:5432/postgres?sslmode=disable -p 8080:8080 goku321/api-server:v0.5

start-server:
	@docker-compose up --detach

stop-server:
	docker-compose down

test:
	go test -count=1 ./... -v

test-e2e:
	@docker-compose up --detach
	go test github.com/goku321/api-server/e2e_test -v
	@docker-compose down

clean:
	go clean

fmt:
	go fmt ./...