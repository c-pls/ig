package resolver

import (
	"database/sql"
	"github.com/c-pls/instagram/backend/db/sqlc"
	"github.com/c-pls/instagram/backend/db/utils"
	_ "github.com/lib/pq"
	"log"
)

// This file will not be regenerated automatically.

// It serves as dependency injection for your app, add any dependencies you require here.

// This file gets initialized once in server.go when we create the graph.

var store *db.Store

//var query *db.Queries

func init() {
	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal(err.Error())
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal(err)
	}
	//query = db.New(conn)
	store = db.NewStore(conn)
}

type Resolver struct {
}
