package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	var saveDataPath string
	if err != nil {
		fmt.Println("Defaulting to saving in current directory")
		saveDataPath = "./expenses.txt"
	} else {
		saveDataPath = os.Getenv("SAVE_DATA_PATH")
		if saveDataPath == "" {
			fmt.Println("SAVE_DATA_PATH not found in environment or .env file")
			fmt.Println("Defaulting to saving in current directory")
			saveDataPath = "./expenses.txt"
		}
	}

	fmt.Println("SAVE_DATA_PATH:", saveDataPath)

	cfg := Config{
		NextID:       0,
		Expenses:     []*Expense{},
		SaveDataPath: saveDataPath,
	}

	if loadData(&cfg) {
		fmt.Println("Loading existing saved data file")
	} else {
		fmt.Println("No existing saved data file - creating new file")
		saveData(&cfg)
	}

	StartREPL(&cfg)
}
