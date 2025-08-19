package controllers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/lenonmerlo/go-currency-api/internal/services"
)

// GetRates godoc
// @Summary Cotação de moedas
// @Description Retorna cotações para os símbolos informados a partir da moeda base.
// @Tags rates
// @Param base query string true "Moeda base" default(BRL)
// @Param symbols query string true "Moedas alvo separadas por vírgula" default(USD,EUR)
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 502 {object} map[string]interface{}
// @Router /v1/rates [get]
func GetRates(c *gin.Context) {
	base := strings.ToUpper(strings.TrimSpace(c.DefaultQuery("base", "BRL")))
	rawSymbols := c.DefaultQuery("symbols", "USD,EUR")

	var symbols []string
	for _, s := range strings.Split(rawSymbols, ",") {
		s = strings.ToUpper(strings.TrimSpace(s))
		if s != "" {
			symbols = append(symbols, s)
		}
	}

	if base == "" || len(symbols) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "bad_request",
			"message": "use ?base=BRL&symbols=USD,EUR (mínimo 1 símbolo)",
		})
		return
	}

	rates, provider, err := services.FetchRates(base, symbols)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"error":   "upstream_error",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"base":     base,
		"rates":    rates,
		"provider": provider,
	})

}
