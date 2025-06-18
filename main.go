package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gunzgo2mars/go-graceful-with-context/src/pkg/cache"
	articleRepo "github.com/gunzgo2mars/go-graceful-with-context/src/repository/article"
	"github.com/gunzgo2mars/go-graceful-with-context/src/service"
	"github.com/redis/go-redis/v9"
)

type HttpServer struct {
	server *http.Server
	svc    service.IService
}

func main() {

	cacheClient := cache.NewRedisConnection(&redis.Options{
		Addr: fmt.Sprintf(
			"%s:%s",
			"localhost",
			"6379",
		),
		DB:          0,
		PoolTimeout: time.Duration(5000),
	})

	articleRepo := articleRepo.New(cacheClient)
	svc := service.New(articleRepo)
	handler := func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		if err := svc.CreateNewArticleInfo(ctx); err != nil {
			http.Error(w, "Failed transaction", http.StatusInternalServerError)
			return
		}

		fmt.Println("Work Done at layer handler")
		w.Write([]byte("Transaction completed"))
	}

	http.HandleFunc("/graceful", handler)
	srv := &HttpServer{
		svc: svc,
		server: &http.Server{
			Addr: ":2318",
		},
	}

	// Channel to listen for termination signals
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	// Run server in goroutine
	go func() {
		if err := srv.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("ListenAndServe(): %v", err)
		}
	}()
	log.Println("Server started")

	// Block until signal received
	<-stop
	log.Println("Shutdown signal received")

	// Create shutdown context with timeout to allow graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	// Shutdown server gracefully
	if err := srv.server.Shutdown(ctx); err != nil {
		log.Fatalf("Server Shutdown Failed:%+v", err)
	}

	log.Println("Server exited properly")

}
