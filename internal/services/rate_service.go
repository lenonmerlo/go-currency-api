package services

import (
	"strings"

	"github.com/lenonmerlo/go-currency-api/internal/clients/exchangerate"
)

func FetchRates(base string, symbols []string) (map[string]float64, error) {

	base = strings.ToUpper(strings.TrimSpace(base))
	var syms []string
	for _, s := range symbols {
		s = strings.ToUpper(strings.TrimSpace(s))
		if s != "" {
			syms = append(syms, s)
		}
	}

	return exchangerate.GetLatest(base, syms)
}
