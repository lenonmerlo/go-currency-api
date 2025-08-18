package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetRates(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H {
		"error": "not_implemented",
		"message": "Rates endpoint will be implemented in the next step.",
	})
}