/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"fmt"
	"log"
	"net/http"
	// "github.com/DaveLaj/budget-tracker/database"
	// "github.com/DaveLaj/budget-tracker/cmd"
	// "github.com/DaveLaj/budget-tracker/utils"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello\n")
}

func main() {
	// cmd.Execute() // command used to enable custom commands

	http.HandleFunc("/hello", hello)

	log.Fatal(http.ListenAndServe("localhost:4200", nil))
}
