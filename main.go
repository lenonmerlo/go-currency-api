package main

import "github.com/gin-gonic/gin"

func main() {
	r:= gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H {
			"message": "Bem-vindo a API de Modeas em GO",
		})
	})

	r.Run(":8080")
}