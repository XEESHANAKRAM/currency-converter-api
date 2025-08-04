// handlers/currency.go - HTTP request handlers
package handlers

import (
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
    "currency-converter/services"
    "currency-converter/models"
)

// ConvertCurrency handles currency conversion requests
func ConvertCurrency(c *gin.Context) {
    // Get query parameters from the request
    from := c.Query("from")
    to := c.Query("to")
    amountStr := c.Query("amount")

    // Validate required parameters
    if from == "" || to == "" || amountStr == "" {
        c.JSON(http.StatusBadRequest, models.ErrorResponse{
            Error: "Missing required parameters: from, to, amount",
        })
        return
    }

    // Convert amount to float
    amount, err := strconv.ParseFloat(amountStr, 64)
    if err != nil {
        c.JSON(http.StatusBadRequest, models.ErrorResponse{
            Error: "Invalid amount format",
        })
        return
    }

    // Get exchange rate from external service
    rate, err := services.GetExchangeRate(from, to)
    if err != nil {
        c.JSON(http.StatusInternalServerError, models.ErrorResponse{
            Error: "Failed to get exchange rate: " + err.Error(),
        })
        return
    }

    // Calculate converted amount
    convertedAmount := amount * rate

    // Return successful response
    c.JSON(http.StatusOK, models.ConversionResponse{
        From: from,
        To: to,
        Amount: amount,
        ConvertedAmount: convertedAmount,
        ExchangeRate: rate,
        Success: true,
    })
}

// GetExchangeRates returns available exchange rates
func GetExchangeRates(c *gin.Context) {
    base := c.DefaultQuery("base", "USD")
    rates, err := services.GetAllRates(base)
    if err != nil {
        c.JSON(http.StatusInternalServerError, models.ErrorResponse{
            Error: err.Error(),
        })
        return
    }

    c.JSON(http.StatusOK, rates)
}

// HealthCheck provides API health status
func HealthCheck(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
        "status": "healthy",
        "message": "Currency Converter API is running",
    })
}