package app

import (
	"log"
	"os"

	"github.com/DaveLaj/budget-tracker/database"
	"github.com/DaveLaj/budget-tracker/internal/utils"
)

func Start() {
	// Main application logic

	file, err := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Error opening/creating log file: %v", err)
	}
	defer file.Close()

	DBConfig, err := utils.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	err = database.InitializeDB(DBConfig, "budget_tracker")
	if err != nil {
		log.Fatalf("Error initializing database: %v", err)
	}

	log.SetOutput(file) // Redirect log output to fil
}
