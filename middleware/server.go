package main

import (
	"log"
	"middleware_loader/graph"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"middleware_loader/middleware/api/auth_service"
)

const defaultPort = "4000"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	authService := auth_service.NewAuthService()

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)
	
	http.HandleFunc("/auth", func(w http.ResponseWriter, r *http.Request) {
		authenticated, err := authService.Auth(r.Context(), r.FormValue("username"), r.FormValue("password"))
	})

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}