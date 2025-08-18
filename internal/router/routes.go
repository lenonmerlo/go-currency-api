package router

import (
	"github.com/gin-gonic/gin"
	"github.com/lenonmerlo/go-currency-api/internal/http/controllers"
)

func Register(r *gin.Engine) {
	// Healthcheck
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	// v1
	v1 := r.Group("/v1")
	{
		v1.GET("/rates", controllers.GetRates)
	}
}
