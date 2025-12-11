package main

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

func saveData(cfg *Config) {

	yamlData, err := yaml.Marshal(&cfg)
	if err != nil {
		log.Fatalf("Error marshaling to YAML: %v", err)
	}

	outputFilePath := cfg.SaveDataPath
	err = os.WriteFile(outputFilePath, yamlData, 0644)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Data saved")
}

func loadData(cfg *Config) bool {
	filePath := cfg.SaveDataPath

	_, err := os.Stat(filePath)
	if err != nil {
		return false
	}

	yamlData, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Error reading save data file: %v", err)
	}

	err = yaml.Unmarshal(yamlData, cfg)
	if err != nil {
		log.Fatal(err)
	}
	return true
}
