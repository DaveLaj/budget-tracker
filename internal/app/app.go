package app

import (
	"log"
	"net/http"
	"os"

	"github.com/DaveLaj/budget-tracker/internal/controllers"
	"github.com/DaveLaj/budget-tracker/internal/database"
	"github.com/DaveLaj/budget-tracker/internal/utils"
	"github.com/gin-gonic/gin"
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

	err = database.InitializeDB(DBConfig, "budget_tracker")
	if err != nil {
		log.Fatalf("Error initializing database: %v", err)
	}

	r := gin.Default()

	r.Static("/static", "./assets")

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("/budgetpage", controllers.ShowPage)

	err = r.Run(":6969")
	if err != nil {
		log.Fatalf("Error running server: %v", err)
	}
}
