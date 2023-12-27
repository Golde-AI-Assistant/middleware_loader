package main

import (
	"log"
	"middleware_loader/graph"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	// "middleware_loader/api"
)

const defaultPort = "4000"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	// authService := api.NewAuthService()

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	http.HandleFunc("/auth", func(w http.ResponseWriter, r *http.Request) {
		username := r.FormValue("username")
		password := r.FormValue("password")

		log.Printf("Username: %s", username)
		log.Printf("Password: %s", password)
	})

	// http.HandleFunc("/auth", func(w http.ResponseWriter, r *http.Request) {
	// 	username := r.FormValue("username")
	// 	password := r.FormValue("password")

	// 	token, err := authService.Auth(r.Context(), username, password)
	// 	log.Printf("token: %s", token)
	// 	if err != nil {
	// 		http.Error(w, err.Error(), http.StatusInternalServerError)
	// 		return
	// 	}
	// })

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
