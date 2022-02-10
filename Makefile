postgres:
	docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=1 -d postgres:12-alpine

createdb:
	docker exec -it postgres12 createdb --username=postgres ig 

dropdb:
	docker exec -it postgres12 dropdb --username=postgres ig 

buildtag:
	go build -tags 'postgres' -ldflags="-X main.Version=1.0.0" -o $GOPATH/bin/migrate github.com/golang-migrate/migrate/v4/cmd/migrate/

migrateup:
	migrate -database "postgres://postgres:1@localhost:5432/ig?sslmode=disable" -path db/migrations -verbose up

migratedown:
	migrate -database "postgres://postgres:1@localhost:5432/ig?sslmode=disable" -path db/migrations -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

graphql:
	 go run github.com/99designs/gqlgen generate

test:
	go test -v -cover ./...

server:
	go run server.go