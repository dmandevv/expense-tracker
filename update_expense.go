package main

import (
	"fmt"
	"strconv"
)

func UpdateExpense(cfg *Config, args ...string) error {
	if len(args) < 6 {
		return fmt.Errorf("not enough arguments, only: %v", len(args))
	}

	id := args[1]
	idNumber, err := strconv.ParseInt(id, 0, 32)
	if err != nil {
		return fmt.Errorf("invalid id format: %v ", id)
	}

	newDescription := args[3]
	newAmount := args[5]
	newAmountNumber, err := strconv.ParseFloat(newAmount, 64)
	if err != nil {
		return fmt.Errorf("invalid new amount: %v ", newAmount)
	}

	var expense *Expense
	for i, e := range cfg.Expenses {
		if e.ID == idNumber {
			expense = cfg.Expenses[i]
			break
		}
	}

	if expense == nil {
		return fmt.Errorf("no expense with ID: %v found", idNumber)
	}

	expense.Description = newDescription
	expense.Amount = newAmountNumber
	fmt.Printf("Expense ID: %v updated\n", idNumber)

	return nil
}
