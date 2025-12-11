package main

func main() {
	cfg := Config{
		NextID:   0,
		Expenses: []Expense{},
	}

	StartREPL(&cfg)

}
