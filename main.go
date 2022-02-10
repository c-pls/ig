package main

import (
	"database/sql"
	"fmt"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/c-pls/instagram/backend/graph/generated"
	"github.com/c-pls/instagram/backend/graph/resolver"
	db2 "github.com/c-pls/instagram/backend/internal/db/sqlc"
	"github.com/c-pls/instagram/backend/internal/db/utils"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

var query *db2.Queries
var store *db2.Store

func main() {
	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(config)
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal(err)
	}
	query = db2.New(conn)
	store = db2.NewStore(conn)
	port := config.Port

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &resolver.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
