package main

import "time"

type Expense struct {
	ID          int64     `yaml:"id"`
	Description string    `yaml:"description"`
	Date        time.Time `yaml:"date"`
	Amount      float64   `yaml:"amount"`
}
