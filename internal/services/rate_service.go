package services

import (
	"fmt"
	"strings"

	"github.com/lenonmerlo/go-currency-api/internal/clients/exchangerate"
	"github.com/lenonmerlo/go-currency-api/internal/clients/frankfurter"
)

func FetchRates(base string, symbols []string) (map[string]float64, string, error) {
	base = strings.ToUpper(strings.TrimSpace(base))

	var syms []string
	for _, s := range symbols {
		s = strings.ToUpper(strings.TrimSpace(s))
		if s != "" {
			syms = append(syms, s)
		}
	}

	if rates, err := exchangerate.GetLatest(base, syms); err == nil && len(rates) > 0 {
		return rates, "exchangerate.host", nil
	}

	if rates, err := frankfurter.GetLatest(base, syms); err == nil && len(rates) > 0 {
		return rates, "frankfurter.app", nil
	}

	return nil, "", fmt.Errorf("no provider returned rates")
}
