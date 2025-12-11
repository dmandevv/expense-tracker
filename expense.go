package main

import "time"

type Expense struct {
	ID          int64
	Description string
	Date        time.Time
	Amount      float64
}
