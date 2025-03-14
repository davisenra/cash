package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
)

const version = "1.0.0"

type CurrencyRates struct {
	Date  string         `json:"date"`
	Rates map[string]any `json:"-"`
}

func main() {
	if len(os.Args) == 2 {
		switch os.Args[1] {
		case "--version":
			fmt.Printf("cash version %s\n", version)
			return
		case "--help":
			fmt.Println("Usage: cash <amount> <from_currency> <to_currency>")
			fmt.Println("Example: cash 10 usd brl")
			fmt.Println("Options:")
			fmt.Println("  --version  Show version information")
			fmt.Println("  --help     Show help information")
			return
		}
	}

	if len(os.Args) != 4 {
		fmt.Println("Error: Invalid number of arguments")
		fmt.Println("Usage: cash <amount> <from_currency> <to_currency>")
		os.Exit(1)
	}

	amount, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil {
		fmt.Println("Error: Invalid amount")
		os.Exit(1)
	}

	fromCurrency := os.Args[2]
	toCurrency := os.Args[3]

	convertedAmount, err := convertCurrency(amount, fromCurrency, toCurrency)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	fmt.Printf("%s $%.2f\n", strings.ToUpper(toCurrency), convertedAmount)
}

func convertCurrency(amount float64, fromCurrency, toCurrency string) (float64, error) {
	url := fmt.Sprintf("https://cdn.jsdelivr.net/npm/@fawazahmed0/currency-api@latest/v1/currencies/%s.json", fromCurrency)
	resp, err := http.Get(url)
	if err != nil {
		return 0, fmt.Errorf("failed to fetch currency rates: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("currency %s not found", fromCurrency)
	}

	var result map[string]any
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return 0, fmt.Errorf("failed to decode response: %v", err)
	}

	rates, ok := result[fromCurrency].(map[string]any)
	if !ok {
		return 0, fmt.Errorf("invalid response format for currency %s", fromCurrency)
	}

	rate, exists := rates[toCurrency]
	if !exists {
		return 0, fmt.Errorf("currency %s not found", toCurrency)
	}

	rateValue, ok := rate.(float64)
	if !ok {
		return 0, fmt.Errorf("invalid rate format for currency %s", toCurrency)
	}

	return amount * rateValue, nil
}
