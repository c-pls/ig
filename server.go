package main

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/c-pls/instagram/backend/db/utils"
	"github.com/c-pls/instagram/backend/graph/generated"
	"github.com/c-pls/instagram/backend/graph/resolver"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

//var query *db2.Queries
//var store *db2.Store

func main() {
	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal(err.Error())
	}

	port := config.Port

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &resolver.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
