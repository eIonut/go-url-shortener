package main

import (
	"context"
	"fmt"
	"go-url-shortener/api"
	"go-url-shortener/config"
	"go-url-shortener/database/db"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/jackc/pgx/v5"
)

func startServer(server *http.Server, serverErr chan<- error) {
	fmt.Println("Server listening on port:", server.Addr)

	err := server.ListenAndServe()

	if err != nil && err != http.ErrServerClosed {
		serverErr <- err
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

	ctx, stop := signal.NotifyContext(
		context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
	)

	defer stop()

	serverErr := make(chan error, 1)

	go startServer(server, serverErr)

	select {
	case <-ctx.Done():
		fmt.Println("shutdown signal received")

	case err := <-serverErr:
		fmt.Println("server error:", err)
		return
	}

	shutdownCtx, cancel := context.WithTimeout(
		context.Background(),
		5*time.Second,
	)

	defer cancel()

	fmt.Println("Server shutting down...")

	err = server.Shutdown(shutdownCtx)
	if err != nil {
		fmt.Println("shutdown error:", err)
	}
}
