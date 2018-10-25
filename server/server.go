package main

import (
	"github.com/s-ichikawa/piql/config"
	"github.com/s-ichikawa/piql/middleware"
	"github.com/s-ichikawa/piql/resolver"
	"log"
	"net/http"

	"github.com/99designs/gqlgen/handler"
	"github.com/s-ichikawa/piql"
)

func main() {
	config, err := config.ReadFromEnv()
	if err != nil {
		log.Fatalf("failed to read env: %s\n", err)
	}

	port := config.PiqlPort

	resolvers := &resolver.Resolver{
		Host: config.PixelaEndpoint,
	}

	http.Handle("/", handler.Playground("GraphQL playground", "/query"))
	http.Handle(
		"/query",
		middleware.AuthMiddleware(handler.GraphQL(piql.NewExecutableSchema(piql.Config{Resolvers: resolvers}))),
	)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
