package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/Salomon-Novachrono/graphQL-test/database"
	"github.com/Salomon-Novachrono/graphQL-test/graph"
	"github.com/Salomon-Novachrono/graphQL-test/graph/generated"
)

const defaultPort = "8080"

func main() {
	database.Connect("mongodb://localhost:27017/")
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))

}
