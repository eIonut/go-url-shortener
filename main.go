package main

import (
	"context"
	"fmt"
	"go-url-shortener/api"
	"go-url-shortener/config"
	"go-url-shortener/database/db"
	"net/http"

	"github.com/jackc/pgx/v5"
)

func startServer(server *http.Server, port string) {
	fmt.Println("Server listening on port:", port)

	err := server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		fmt.Println("Server error:", err)
	}
}

func main() {
	cfg := config.Load()

	conn, err := pgx.Connect(context.Background(), cfg.DatabaseURL)
	if err != nil {
		panic(err)
	}
	defer conn.Close(context.Background())

	queries := db.New(conn)

	handler := api.NewHandler(queries)
	router := api.NewRouter(handler)

	server := &http.Server{
		Addr:    ":" + cfg.Port,
		Handler: router,
	}
	startServer(server, cfg.Port)
}
