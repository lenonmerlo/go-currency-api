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
	Base  string             `json:"base"`
	Rates map[string]float64 `json:"rates"`
}

var httpClient = &http.Client{Timeout: 8 * time.Second}


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
		return nil, fmt.Errorf("provider status %d", res.StatusCode)
	}

	var payload latestResp
	if err := json.NewDecoder(res.Body).Decode(&payload); err != nil {
		return nil, err
	}

	return payload.Rates, nil
}
