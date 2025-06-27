
package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"url-shortener/internal/handler"
)

func main() {
	r := gin.Default()
	
	handler.RegisterRouters(r)

	if err := r.Run(":8000"); err != nil {
		log.Fatal(err)
	}
}