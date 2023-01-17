package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/teste-unico/db"
	"github.com/teste-unico/handler"
)

func main() {
	port := "localhost:8080"
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Error occured: %s", err.Error())
	}

	dbUser := "unico"
	dbPassword := "unico123"
	dbName := "unico"

	database, err := db.Initialize(dbUser, dbPassword, dbName)
	if err != nil {
		log.Fatalf("Could not set up database: %v", err)
	}

	defer database.Conn.Close()

	httpHandler := handler.NewHandler(database)
	server := &http.Server{
		Handler: httpHandler,
	}

	go func() {
		server.Serve(listener)
	}()

	defer Stop(server)

	log.Printf("Started server on %s", port)
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	log.Println(fmt.Sprint(<-ch))
	log.Println("Stopping API server.")
}

func Stop(server *http.Server) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Printf("Could not shut down server correctly: %v\n", err)
		os.Exit(1)
	}
}
