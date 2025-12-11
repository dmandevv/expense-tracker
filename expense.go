package main

import "time"

type Expense struct {
	ID     int
	Name   string
	Date   time.Time
	Amount float64
}
