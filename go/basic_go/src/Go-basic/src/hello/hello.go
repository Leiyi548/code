package main

import (
	"example.com/greetings"
	"fmt"
	"log"
)

func main() {
	// set properties of the predefined Logger,including
	// the log entry prefix and a flag to disable printing
	// the time,source file,and line number
	log.SetPrefix("greetings: ")
	log.SetFlags(log.Llongfile | log.Ltime | log.Ldate | log.Lmicroseconds)
	fmt.Println(log.Flags())

	// A slice of name
	names := []string{"Gladys", "Samntha", "Darrin"}

	// Request greeting messages for the names.
	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}

	// if no error was returned, print the returned map of
	fmt.Println(messages)

	// Request  a greeting message.
	message, err := greetings.Hello("Leiyi548")
	// if an error was returned,print it to the console and
	// exit the program.
	if err != nil {
		log.Fatal(err)
	}
	// If no error was returned,print the returned message
	// to the console.
	fmt.Println(message)
}
