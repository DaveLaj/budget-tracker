/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"fmt"

	"github.com/DaveLaj/dateExercise/database"
	// "github.com/DaveLaj/dateExercise/cmd"
)

func main() {
	// cmd.Execute() // command used to enable custom commands
	cfg, err := LoadConfig()
	if err != nil {
		fmt.Println("Error loading config")
		return
	}
	db, err := database.CreateDBConnection(cfg, "budget")
	if err != nil {
		fmt.Println("Error connecting to database")
		return
	}
	defer db.Close()

}
