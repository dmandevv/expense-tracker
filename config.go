package main

type Config struct {
	NextID       int64      `yaml:"next_id"`
	Expenses     []*Expense `yaml:"expenses"`
	SaveDataPath string     `yaml:"-"`
}
