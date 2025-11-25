package main

import (
	"log"
	"net"

	"os"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/Shobhit150/url_shortner/internal/handler"
	"github.com/Shobhit150/url_shortner/internal/kafka"
	"github.com/Shobhit150/url_shortner/internal/middleware"
	"github.com/Shobhit150/url_shortner/internal/repository"
	urlshortenerpb "github.com/Shobhit150/url_shortner/proto"
)

func main() {
	dsn := os.Getenv("DATABASE_URL")
    if dsn == "" {
        // Use `db` as the hostname when running in Docker Compose!
        dsn = "postgres://user:password@db:5432/urlshortener?sslmode=disable"
    }
  

	kafka.InitKafka()
	repository.InitPostgres(dsn)

	go kafka.ReadFromKafka()
	
	// REST API
	go func() {
		r := gin.Default()
		

		r.Use(func(c *gin.Context) {
			c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
			c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Accept, Authorization")
			if c.Request.Method == "OPTIONS" {
				c.AbortWithStatus(204)
				return
			}
			c.Next()
		})
		
		r.Use(middleware.RateLimiter())
		handler.RegisterRouters(r)
			
		log.Println("REST API listening on the :8080")
		if err := r.Run(":8080"); err != nil {
			log.Fatal("REST server:", err)
		}
	}()
	

	
	// gRPC API
	go func() {
		lis, err := net.Listen("tcp", ":50051")
		if err != nil {
			log.Fatal("failed to listen:", err)
		}
		grpcServer := grpc.NewServer()
		urlshortenerpb.RegisterURLShortenerServer(grpcServer, &handler.URLShortenerServer{})

		reflection.Register(grpcServer) 

		log.Println("gRPC API listening on :50051")
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatal("gRPC server:", err)
		}
	}()

	

	select {} // Block forever
}
