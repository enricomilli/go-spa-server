package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/enricomilli/go-spa-server/ui"
	"github.com/go-chi/chi/v5"
)

func main() {
	if err := run(); err != nil {
		log.Fatalln(err)
	}
}

func run() error {

	router := chi.NewRouter()

	ui.SetupRoutes(router)

	port := setupPort()

	fmt.Println("Server started on port", port)
	return http.ListenAndServe(port, router)
}

func setupPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	return ":" + port
}
