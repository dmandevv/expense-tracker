package main

import (
	"fmt"
	"slices"
	"strconv"
)

func DeleteExpense(cfg *Config, args ...string) error {
	if len(args) < 2 {
		return fmt.Errorf("not enough arguments, only: %v", len(args))
	}

	id := args[1]
	idNumber, err := strconv.ParseInt(id, 0, 32)
	if err != nil {
		return fmt.Errorf("invalid id format: %v ", id)
	}

	for i, e := range cfg.Expenses {
		if e.ID == idNumber {
			cfg.Expenses = slices.Delete(cfg.Expenses, i, i+1)
			fmt.Printf("Expense ID: %v deleted\n", idNumber)

			saveData(cfg)

			return nil
		}
	}
	return fmt.Errorf("no expense with ID: %v found", idNumber)
}
