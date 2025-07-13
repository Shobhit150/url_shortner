package main

import (
	"log"
	"os"

	"github.com/Shobhit150/url_shortner/internal/handler"
	"github.com/Shobhit150/url_shortner/internal/repository"
	"github.com/gin-gonic/gin"
)

func main() {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		dsn = "postgres://user:password@localhost:5432/urlshortener?sslmode=disable"
	}

	repository.InitPostgres(dsn)

	r := gin.Default()
	
	handler.RegisterRouters(r)

	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}