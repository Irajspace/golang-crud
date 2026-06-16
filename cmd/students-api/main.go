package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"log"
	"github.com/irajspace/golang-crud/internal/config"
	"github.com/irajspace/golang-crud/internal/http/handlers/student"
	"github.com/irajspace/golang-crud/internal/storage/sqlite"
)

func main() {
	// Load config
	cfg := config.MustLoad()
	storage, err := sqlite.New(cfg)
	if err != nil {
		log.Fatal("failed to initialize storage", "error", err)
	}
	slog.Info("config loaded successfully",slog.String("env", cfg.Env), slog.String("storage_path", cfg.StoragePath), slog.String("http_addr", cfg.HTTPServer.Addr))
	router := http.NewServeMux()

	router.HandleFunc("POST /api/students", student.New(storage))

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

	error := server.Shutdown(ctx)

	if error != nil {
		slog.Error("error shutting down server", "error", err)
	}
	slog.Info("server gracefully stopped")
}