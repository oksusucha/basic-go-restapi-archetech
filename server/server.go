package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var r gin.Engine

func init() {
	if err := godotenv.Load("server/.env"); err != nil {
		log.Fatal(err)
	}

	initializeLogging()
}

func Run() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	r = *gin.Default()
	initRouter()

	srv := &http.Server{
		Addr: func() string {
			port := os.Getenv("RUN_PORT")
			if port == "" {
				port = "8080"
			}

			return fmt.Sprintf(":%s", port)
		}(),
		Handler: &r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	<-ctx.Done()

	stop()
	log.Println("Shutting down gracefully, press Ctrl+C again to force")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown: ", err)
	}

	log.Println("Server exiting")
}
