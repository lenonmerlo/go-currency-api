package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/lenonmerlo/go-currency-api/internal/router"
)

func main() {
	r:= gin.Default()
	router.Register(r)

	log.Println("listening on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}