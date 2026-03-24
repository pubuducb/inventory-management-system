package main

import (
	"context"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"ims/internal/config"
	"ims/internal/server"
)

func main() {

	db, err := config.InitDB()
	if err != nil {
		log.Fatalf("[DB] Failed to initialize database: %v", err)
	}
	defer db.Close()

	if err := config.SetupSchema(db); err != nil {
		log.Fatalf("[DB] Failed to setup schema: %v", err)
	}

	httpServer := server.InitHTTPServer(db)
	go startServer(httpServer)
	waitForShutdown(httpServer)
}

func startServer(server *http.Server) {
	log.Printf("[Server] Starting on port %s", server.Addr)

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("[Server] Listen: %s\n", err)
	}
}

func waitForShutdown(server *http.Server) {

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()
	<-ctx.Done()
	log.Printf("[Server] Shutting down")

	ctx, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFunc()

	if err := server.Shutdown(ctx); err != nil {
		log.Printf("[Server] Forced shutdown: %v", err)
	}

	log.Printf("[Server] Exiting properly")
}
