package main

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/c-pls/instagram/backend/db/utils"
	"github.com/c-pls/instagram/backend/graph/generated"
	"github.com/c-pls/instagram/backend/graph/resolver"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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

	e := echo.New()

	// Middleware
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(middleware.Gzip())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	e.GET("/health", func(c echo.Context) error {
		return c.NoContent(http.StatusOK)
	})

	//srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &resolver.Resolver{}}))

	e.POST("/query", func(c echo.Context) error {
		srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &resolver.Resolver{}}))

		srv.ServeHTTP(c.Response(), c.Request())

		return nil
	})

	e.GET("/", func(c echo.Context) error {
		h := playground.Handler("Test", "/query")
		h.ServeHTTP(c.Response(), c.Request())
		return nil

	})

	//http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	//http.Handle("/query", srv)

	e.HideBanner = true
	e.Logger.Fatal(e.Start(":" + port))

	//router := chi.NewRouter()
	//
	//router.Use(cors.New(cors.Options{
	//	AllowedOrigins:   []string{"http://localhost:3000"},
	//	AllowCredentials: true,
	//	Debug:            true,
	//}).Handler)
	//
	//
	//srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &resolver.Resolver{}}))
	//srv.AddTransport(&transport.Websocket{
	//	Upgrader: websocket.Upgrader{
	//		CheckOrigin: func(r *http.Request) bool {
	//			return r.Host == "http://localhost:3000/"
	//		},
	//		ReadBufferSize:  1024,
	//		WriteBufferSize: 1024,
	//	},
	//})
	//
	//router.Handle("/", playground.Handler("Graphql Playground", "/query"))
	//router.Handle("/query", srv)
	//
	//err = http.ListenAndServe(":"+port, router)
	//if err != nil {
	//	panic(err)
	//}

}
