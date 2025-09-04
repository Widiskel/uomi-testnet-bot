package config

import (
	"encoding/json"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	AccountsPath string
	TxAmountMin  float64
	TxAmountMax  float64
}

func Load() Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using default values")
	}

	minStr := os.Getenv("TX_AMOUNT_MIN")
	min, err := strconv.ParseFloat(minStr, 64)
	if err != nil || min <= 0 {
		min = 0.01
	}

	maxStr := os.Getenv("TX_AMOUNT_MAX")
	max, err := strconv.ParseFloat(maxStr, 64)
	if err != nil || max <= min {
		max = 0.05
	}

	return Config{
		AccountsPath: "configs/accounts.json",
		TxAmountMin:  min,
		TxAmountMax:  max,
	}
}

func (c Config) LoadAccounts() ([]string, error) {
	b, err := os.ReadFile(c.AccountsPath)
	if err != nil {
		return nil, err
	}
	var accounts []string
	return accounts, json.Unmarshal(b, &accounts)
}
