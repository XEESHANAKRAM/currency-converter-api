// models/response.go - API response structures
package models

// ConversionResponse represents a successful currency conversion
type ConversionResponse struct {
    From string `json:"from"`
    To string `json:"to"`
    Amount float64 `json:"amount"`
    ConvertedAmount float64 `json:"converted_amount"`
    ExchangeRate float64 `json:"exchange_rate"`
    Success bool `json:"success"`
}

// ErrorResponse represents an API error
type ErrorResponse struct {
    Error string `json:"error"`
    Success bool `json:"success" default:"false"`
}

// ExchangeRateResponse represents exchange rate data
type ExchangeRateResponse struct {
    Base string `json:"base"`
    Rates map[string]float64 `json:"rates"`
    Success bool `json:"success"`
}