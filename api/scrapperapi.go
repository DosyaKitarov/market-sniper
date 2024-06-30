package api

import (
	"context"
	"fmt"
	"github.com/DosyaKitarov/market-sniper/internal/pkg/env"
	"io"
	"net/http"
	"net/url"
	"time"
)

const (
	TimeOut = 10 * time.Second
	baseURL = "https://api.scraperapi.com/structured/amazon/product"
)

var (
	apiKey = env.GetEnvVariable("API_KEY_SCRAPPER")
)

func FetchProductData(ctx context.Context, asin, countryCode, tld string) (string, error) {
	params := url.Values{}
	params.Add("api_key", apiKey)
	params.Add("asin", asin)
	params.Add("country_code", countryCode)
	params.Add("tld", tld)

	req, err := http.NewRequestWithContext(ctx, "GET", baseURL+"?"+params.Encode(), nil)
	if err != nil {
		return "", fmt.Errorf("error creating request: %w", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("error fetching data: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("API request failed with status: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response body: %w", err)
	}

	return string(body), nil
}
