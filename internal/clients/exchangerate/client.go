package exchangerate

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type latestResp struct {
	Success *bool              `json:"success"` // pode existir
	Base    string             `json:"base"`
	Rates   map[string]float64 `json:"rates"`
	Error   any                `json:"error"`   // pode vir em alguns planos
}

var httpClient = &http.Client{Timeout: 8 * time.Second}

// https://api.exchangerate.host/latest?base=BRL&symbols=USD,EUR
func GetLatest(base string, symbols []string) (map[string]float64, error) {
	if base == "" || len(symbols) == 0 {
		return nil, fmt.Errorf("invalid params: base=%q symbols=%v", base, symbols)
	}

	apiURL := "https://api.exchangerate.host/latest"
	q := url.Values{}
	q.Set("base", strings.ToUpper(base))
	q.Set("symbols", strings.ToUpper(strings.Join(symbols, ",")))

	req, err := http.NewRequest(http.MethodGet, apiURL+"?"+q.Encode(), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")

	res, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode < 200 || res.StatusCode >= 300 {
		return nil, fmt.Errorf("exchangerate status %d", res.StatusCode)
	}

	var payload latestResp
	if err := json.NewDecoder(res.Body).Decode(&payload); err != nil {
		return nil, err
	}

	// se vier success=false ou rates vazio, trate como erro
	if payload.Success != nil && !*payload.Success {
		return nil, fmt.Errorf("exchangerate: success=false (error=%v)", payload.Error)
	}
	if len(payload.Rates) == 0 {
		return nil, fmt.Errorf("exchangerate: empty rates")
	}

	return payload.Rates, nil
}
