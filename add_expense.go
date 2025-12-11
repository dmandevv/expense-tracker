package main

import (
	"fmt"
	"strconv"
	"time"
)

func AddExpense(cfg *Config, args ...string) error {
	if len(args) < 4 {
		return fmt.Errorf("not enough arguments, only: %v", len(args))
	}

	description := args[1]
	amount := args[3]
	amountFloat, err := strconv.ParseFloat(amount, 64)
	if err != nil {
		return fmt.Errorf("invalid amount format: %v", amount)
	}

	newExpense := Expense{
		ID:          cfg.NextID,
		Description: description,
		Date:        time.Now().UTC(),
		Amount:      amountFloat,
	}
	cfg.NextID += 1
	cfg.Expenses = append(cfg.Expenses, &newExpense)

	fmt.Printf("Added expense (ID: %v)\n", newExpense.ID)

	return nil
}
