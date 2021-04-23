package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	// Set properties of the predefined Logger, including
	// the log
	log.SetPrefix("greetings:")
	log.SetFlags(0)
	// Get a greeting message and print it.
	//message, err := greetings.Hello("Gladys")
	//message , err:= greetings.Hello("")
	names := []string{
		"Darrin",
		"Gladys",
		"Samantha",
	}
	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(messages)
}
