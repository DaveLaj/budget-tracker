/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"fmt"
	// "github.com/DaveLaj/dateExercise/cmd"
)

func main() {
	// cmd.Execute()
	cfg, err := LoadConfig()
	if err != nil {
		fmt.Println("Error loading config")
		return
	}
	CreateDBConnection(cfg, "budget")
}
