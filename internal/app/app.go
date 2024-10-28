package app

import (
	"log"
	"os"

	"github.com/DaveLaj/budget-tracker/internal/utils"
)

func Start() {
	// Main application logic

	file, err := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Error opening/creating log file: %v", err)
	}
	defer file.Close()

	log.SetOutput(file) // Redirect log output to file

	DBConfig, err := utils.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	log.Printf("Config: %+v", DBConfig)
}
