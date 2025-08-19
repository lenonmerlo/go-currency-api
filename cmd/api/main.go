package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/lenonmerlo/go-currency-api/internal/router"

	docs "github.com/lenonmerlo/go-currency-api/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Go Currency API
// @version 1.0
// @description API em Go para cotações de moedas (com fallback de provider).
// @host localhost:8080
// @BasePath /
func main() {
	r := gin.Default()
	r.SetTrustedProxies(nil)

	// rotas da aplicação
	router.Register(r)

	// configurações do swagger
	docs.SwaggerInfo.Title = "Go Currency API"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.BasePath = "/"

	// endpoint da UI
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// porta via env
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Println("listening on :" + port)
	if err := r.Run(":" + port); err != nil {
		log.Fatal(err)
	}
}
