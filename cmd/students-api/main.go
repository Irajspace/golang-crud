package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/irajspace/golang-crud/internal/config"
	"github.com/irajspace/golang-crud/internal/http/handlers/student"
)

func main() {
	// Load config
	cfg := config.MustLoad()

	router := http.NewServeMux()

	router.HandleFunc("POST /api/students", student.New())

	server := &http.Server{
		Addr:    cfg.HTTPServer.Addr,
		
		Handler: router,
	}
	slog.Info("starting server", "address", cfg.HTTPServer.Addr)
	done :=make(chan os.Signal,1)

	signal.Notify(done, os.Interrupt,syscall.SIGINT, syscall.SIGTERM)
	go func(){
		err := server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			panic(err)
		}
	}()
	
	<-done

	slog.Info("shutting down server")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := server.Shutdown(ctx)

	if err != nil {
		slog.Error("error shutting down server", "error", err)
	}
	slog.Info("server gracefully stopped")
}