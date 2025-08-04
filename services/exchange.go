// services/exchange.go - External API integration
package services

import (
    "encoding/json"
    "fmt"
    "net/http"
    "os"
    "time"
)

const (
    exchangeAPIURL = "https://api.exchangerate-api.com/v4/latest/"
    timeout = 10 * time.Second
)

// ExchangeAPIResponse represents the external API response
type ExchangeAPIResponse struct {
    Base string `json:"base"`
    Rates map[string]float64 `json:"rates"`
}

// GetExchangeRate fetches conversion rate between two currencies
func GetExchangeRate(from, to string) (float64, error) {
    // Create HTTP client with timeout
    client := &http.Client{
        Timeout: timeout,
    }

    // Build API URL
    url := exchangeAPIURL + from
    if apiKey := os.Getenv("EXCHANGE_API_KEY"); apiKey != "" {
        url += "?access_key=" + apiKey
    }

    // Make HTTP request
    resp, err := client.Get(url)
    if err != nil {
        return 0, fmt.Errorf("failed to fetch exchange rate: %w", err)
    }
    defer resp.Body.Close()

    // Check HTTP status
    if resp.StatusCode != http.StatusOK {
        return 0, fmt.Errorf("API returned status %d", resp.StatusCode)
    }

    // Parse JSON response
    var result ExchangeAPIResponse
    if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
        return 0, fmt.Errorf("failed to parse API response: %w", err)
    }

    // Get conversion rate
    rate, exists := result.Rates[to]
    if !exists {
        return 0, fmt.Errorf("currency %s not found", to)
    }

    return rate, nil
}

// GetAllRates fetches all available exchange rates for a base currency
func GetAllRates(base string) (*ExchangeAPIResponse, error) {
    client := &http.Client{Timeout: timeout}
    url := exchangeAPIURL + base

    resp, err := client.Get(url)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    var result ExchangeAPIResponse
    err = json.NewDecoder(resp.Body).Decode(&result)
    return &result, err
}