package server

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var r *gin.Engine

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	// TODO: setting logrus
}

func Run() {
	r := *gin.Default()

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
}
