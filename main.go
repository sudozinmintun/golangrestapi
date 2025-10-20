package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/sudozinmintun/golangrestapi/dbconfig"
	"github.com/sudozinmintun/golangrestapi/internal/handlers"
	"github.com/sudozinmintun/golangrestapi/internal/routes"
	"github.com/sudozinmintun/golangrestapi/internal/store"
	"github.com/sudozinmintun/golangrestapi/serverconfig"
)

func main() {
	config, err := serverconfig.LoadConfig()

	if err != nil {
		log.Fatalf("Could not load config: %v", err)
	}

	db := dbconfig.ConnectDB(config.DatabaseURL)
	defer db.Close()

	queries := store.New(db)
	handler := handlers.NewHandler(db, queries)

	mux := http.NewServeMux()

	routes.SetupRoutes(mux, handler)

	serverAddr := fmt.Sprintf(":%s", config.ServerPort)
	server := &http.Server{
		Addr:    serverAddr,
		Handler: mux,
	}

	fmt.Printf("Server is up and running on PORT %s", serverAddr)
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
