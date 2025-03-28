package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Exchange rates (hardcoded for simplicity)
var exchangeRates = map[string]float64{
	"USD": 1.0,
	"EUR": 0.91,
	"GBP": 0.78,
	"AUD": 1.52,
	"JPY": 110.25,
}

// Response structure
type ConversionResult struct {
	From   string  `json:"from"`
	To     string  `json:"to"`
	Amount float64 `json:"amount"`
	Result float64 `json:"result"`
}

// Currency conversion handler
func convertCurrency(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	from := query.Get("from")
	to := query.Get("to")
	amount := query.Get("amount")

	// Validate query parameters
	var value float64
	_, err := fmt.Sscanf(amount, "%f", &value)
	if err != nil || exchangeRates[from] == 0 || exchangeRates[to] == 0 {
		http.Error(w, "Invalid parameters", http.StatusBadRequest)
		return
	}

	// Convert currency
	result := (value / exchangeRates[from]) * exchangeRates[to]

	// Respond with JSON
	response := ConversionResult{From: from, To: to, Amount: value, Result: result}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/convert", convertCurrency).Methods("GET")

	fmt.Println("Currency Converter API is running on port 8080...")
	http.ListenAndServe(":8080", r)
}

