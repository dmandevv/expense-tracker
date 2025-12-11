package main

import (
	"fmt"
	"os"
	"text/tabwriter"
)

func ListExpenses(cfg *Config, args ...string) error {
	tabWriter := tabwriter.NewWriter(os.Stdout, 0, 0, 5, ' ', 0)
	defer tabWriter.Flush()
	fmt.Fprintln(tabWriter, "ID\tDate\tDescription\tAmount")
	for _, e := range cfg.Expenses {
		fmt.Fprintf(tabWriter, "%v\t%v\t%v\t$%.2f\n", e.ID, e.Date.Local().Format("2006-01-02"), e.Description, e.Amount)
	}
	return nil
}
