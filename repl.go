package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type command struct {
	description string
	usage       string
	callback    func(*Config, ...string) error
}

func StartREPL(cfg *Config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		print("Expense Tracker > ")
		scanner.Scan()
		cleanedInput := cleanInput(scanner.Text())
		if len(cleanedInput) < 1 {
			continue
		}
		command, exists := getCommands()[cleanedInput[0]]
		if exists {
			args := cleanedInput[1:]
			err := command.callback(cfg, args...)
			if err != nil {
				fmt.Println("Error:", err)
			}
			continue
		} else {
			fmt.Println("Unknown Command")
		}
	}

}

func cleanInput(input string) []string {
	lowerInput := strings.ToLower(input)
	noQuotes := strings.ReplaceAll(lowerInput, `"`, ``)
	trimmedInput := strings.Fields(noQuotes)
	return trimmedInput
}

func getCommands() map[string]command {
	return map[string]command{
		"add": {
			description: "Add a new expense",
			usage:       "add --description <\"description\"> --amount <0.00>",
			callback:    AddExpense,
		},
		"list": {
			description: "List all expenses",
			usage:       "list",
			callback:    ListExpenses,
		},
		"summary": {
			description: "List total amount of expenses",
			usage:       "summary [--month <number>]",
			callback:    Summary,
		},
		"update": {
			description: "Change an expense's info based on ID",
			usage:       "update --id <id> --description <\"description\"> --amount <0.00>",
			callback:    UpdateExpense,
		},
		"delete": {
			description: "Delete an expense by its ID",
			usage:       "delete --id <id>",
			callback:    DeleteExpense,
		},
		"help": {
			description: "List available commands",
			usage:       "help",
			callback:    help,
		},
		"exit": {
			description: "Exit program",
			usage:       "exit",
			callback:    exit,
		},
	}
}

func help(cfg *Config, args ...string) error {
	mapOfCommands := getCommands()
	fmt.Print("--List of commands--")
	for name, command := range mapOfCommands {
		fmt.Printf("\n\n%v: %v", name, command.description)
		fmt.Printf("\nUsage: %v", command.usage)
	}
	fmt.Println()
	fmt.Println()
	return nil
}

func exit(cfg *Config, args ...string) error {
	os.Exit(0)
	return nil
}
