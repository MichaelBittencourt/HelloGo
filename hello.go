package main

import (
	"fmt"
	"log"

	"github.com/MichaelBittencourt/greatings"
	"rsc.io/quote"
)

func main() {
	fmt.Println(quote.Go())

	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number
	log.SetPrefix("greatings: ")
	log.SetFlags(0)

	// Request a greatings message.
	message, err := greatings.Hello("Michael")
	// If an error was returned, print it to the console and
	// exit the program
	if err != nil {
		log.Fatal(err)
	}

	// If no error was returned, print the returned message
	// to the console
	fmt.Println(message)
}
