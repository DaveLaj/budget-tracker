/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/DaveLaj/budget-tracker/database"
	"github.com/DaveLaj/budget-tracker/utils"
	"github.com/spf13/cobra"
)

// createCmd represents the create command
var sqlCmd = &cobra.Command{
	Use:   "sql",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Ongoing work on sql command.")

		// Setup for the LoadConfig function
		cfg, err := utils.LoadConfig()
		if err != nil {
			fmt.Println("Error loading config")
			return
		}
		db, err := database.CreateDBConnection(cfg, "budget")
		if err != nil {
			fmt.Println("Error connecting to database")
		}
		defer db.Close()

		err = db.Ping()
		if err != nil {
			fmt.Println("Error pinging the database")
			return
		}
		// This is where the code for the sql command will go
		// create table example
		table, err := db.Query("CREATE TABLE IF NOT EXISTS expenses (id SERIAL PRIMARY KEY, name TEXT, amount FLOAT, date DATE);")
		if err != nil {
			fmt.Println("Error creating table")
		} else {
			fmt.Println("Table created successfully")
		}
		defer table.Close()

		// build := exec.Command("go", "build", "-o", "sql", "sql.go")
		// build.run() // run the build command
	},
}

func init() {
	createCmd.AddCommand(sqlCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
