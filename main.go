package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	saveDataPath := os.Getenv("SAVE_DATA_PATH")
	if saveDataPath == "" {
		fmt.Println("SAVE_DATA_PATH not found in environment or .env file")
		fmt.Println("Defaulting to current directory")
		saveDataPath = "./expenses.txt"
	} else {
		fmt.Println("SAVE_DATA_PATH:", saveDataPath)
	}

	cfg := Config{
		NextID:       0,
		Expenses:     []*Expense{},
		SaveDataPath: saveDataPath,
	}

	if loadData(&cfg) {
		fmt.Println("Loading existing saved expenses data")
	} else {
		fmt.Println("No existing saved expenses data: Creating new file")
		saveData(&cfg)
	}

	StartREPL(&cfg)

}
