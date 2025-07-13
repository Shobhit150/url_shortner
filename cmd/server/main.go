package main

import (
	"log"
	"net"
	"os"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection" // ADD THIS LINE

	"github.com/Shobhit150/url_shortner/internal/handler"
	"github.com/Shobhit150/url_shortner/internal/repository"
	urlshortenerpb "github.com/Shobhit150/url_shortner/proto"
)

func main() {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		dsn = "postgres://user:password@localhost:5432/urlshortener?sslmode=disable"
	}
	repository.InitPostgres(dsn)

	// REST API
	go func() {
		r := gin.Default()
		handler.RegisterRouters(r)
		log.Println("REST API listening on :8080")
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

		reflection.Register(grpcServer) // <-- ADD THIS

		log.Println("gRPC API listening on :50051")
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatal("gRPC server:", err)
		}
	}()

	select {} // Block forever
}
