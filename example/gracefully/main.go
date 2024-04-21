package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	mux := http.NewServeMux()

	srv := http.Server{
		Addr:              ":" + os.Getenv("APP_PORT"),
		Handler:           mux,
		ReadHeaderTimeout: 5 * time.Second,
	}

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	idleConnsClosed := make(chan struct{})

	go func() {
		<-ctx.Done()

		d := time.Duration(5 * time.Second)
		fmt.Printf("shutting down int %s ...", d)
		// We received an interrupt signal, shut down.
		ctx, cancel := context.WithTimeout(context.Background(), d)
		defer cancel()
		if err := srv.Shutdown(ctx); err != nil {
			// Error from closing listeners, or context timeout:
			slog.Info("HTTP server Shutdown: " + err.Error())
		}
		close(idleConnsClosed)
	}()

	fmt.Println(":" + os.Getenv("APP_PORT") + " is serve")

	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		slog.Error("HTTP server ListenAndServe: " + err.Error())
		return
	}

	<-idleConnsClosed
	fmt.Println("bye")
}
