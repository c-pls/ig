postgres:
	docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=1 -d postgres:12-alpine

create-db:
	docker exec -it postgres12 createdb --username=postgres ig 

drop-db:
	docker exec -it postgres12 dropdb --username=postgres ig 

build-tag:
	go build -tags 'postgres' -ldflags="-X main.Version=1.0.0" -o $GOPATH/bin/migrate github.com/golang-migrate/migrate/v4/cmd/migrate/

migrate-up:
	migrate -database "postgres://postgres:1@localhost:5432/ig?sslmode=disable" -path db/migrations -verbose up

migrate-down:
	migrate -database "postgres://postgres:1@localhost:5432/ig?sslmode=disable" -path db/migrations -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

graphql-gen:
	 go run github.com/99designs/gqlgen generate

server-start:
	go run server.go

test:
	go test -v -cover ./...
