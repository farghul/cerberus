package main

import (
	"fmt"
	"log"
	"os"
)

type Color string

const (
	Reset         = "\033[0m"
	Red    Color  = "\033[31m"
	Green  Color  = "\033[32m"
	Yellow Color  = "\033[33m"
	BGRed  Color  = "\033[41m"
	bv     string = "1.0.0"
	halt   string = "program halted "
)

// Provide and highlight an informational message
func inform(message string) {
	Yellow.Printf("%s", "** ")
	fmt.Print(message)
	Yellow.Println(" **")
}

// Print a colourized error message
func alert(message string) {
	Red.Printf("\n%s", "Error: ")
	fmt.Printf("%s", message)
	BGRed.Println(halt)
	inform("Use -help to display help information")
	os.Exit(0)
}

// Check for errors, log the result if found
func inspect(err error) {
	if err != nil {
		log.Fatal(err)
		return
	}
}

// Println function for colourized text
func (c Color) Println(text string) {
	fmt.Println(string(c) + text + Reset)
}

// Printf function for colourized text
func (c Color) Printf(format string, a ...any) {
	fmt.Printf(string(c)+format+Reset, a...)
}
