package frankfurter

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type latestResp struct {
	Base  string             `json:"base"`
	Rates map[string]float64 `json:"rates"`
}

var httpClient = &http.Client{Timeout: 8 * time.Second}

// https://api.frankfurter.app/latest?from=BRL&to=USD,EUR
func GetLatest(from string, to []string) (map[string]float64, error) {
	if from == "" || len(to) == 0 {
		return nil, fmt.Errorf("invalid params: from=%q to=%v", from, to)
	}

	apiURL := "https://api.frankfurter.app/latest"
	q := url.Values{}
	q.Set("from", strings.ToUpper(from))
	q.Set("to", strings.ToUpper(strings.Join(to, ",")))

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
		return nil, fmt.Errorf("frankfurter status %d", res.StatusCode)
	}

	var payload latestResp
	if err := json.NewDecoder(res.Body).Decode(&payload); err != nil {
		return nil, err
	}
	if len(payload.Rates) == 0 {
		return nil, fmt.Errorf("frankfurter: empty rates")
	}
	return payload.Rates, nil
}
