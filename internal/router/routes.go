package router

import (
	"github.com/lenonmerlo/go-currency-api/internal/http/controllers"
	"github.com/gin-gonic/gin"

)

func Register(r *gin.Engine) {
	r.GET("/health", func (c *gin.Context)  {
		c.JSON(200, gin.H{"status": "ok"})
	})

	v1 := r.Group("/v1")
	{
		v1.GET("/rates", controllers.GetRates)
	}
}