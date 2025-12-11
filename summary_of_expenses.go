package main

import (
	"fmt"
	"strconv"
	"time"
)

func Summary(cfg *Config, args ...string) error {
	total := 0.00
	if len(args) > 1 { //month specified
		monthNumber, err := strconv.ParseInt(args[1], 0, 8)
		month := time.Month(monthNumber)
		if err != nil || month < 1 || month > 12 {
			return fmt.Errorf("invalid month format: %v", args[1])
		}
		for _, e := range cfg.Expenses {
			if e.Date.Month() == month {
				total += e.Amount
			}
		}
		fmt.Printf("Total expenses for %v: $%.2f\n", month.String(), total)
		return nil
	}
	for _, e := range cfg.Expenses {
		total += e.Amount
	}
	fmt.Printf("Total expenses: $%.2f\n", total)
	return nil
}
