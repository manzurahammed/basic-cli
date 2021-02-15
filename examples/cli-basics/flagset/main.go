package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	osParam := os.Args
	if len(osParam)<2 {
		fmt.Println("No command provided")
		os.Exit(2)
	}

	cmd := osParam[1]

	switch cmd {
	case "greet":
		greetMd := flag.NewFlagSet("greet",flag.ExitOnError)
		message := greetMd.String("msg", "Basic CLI Command","The message");
		err := greetMd.Parse(osParam[2:])
		if err !=nil {
			log.Fatal(err.Error())
		}
		fmt.Println("Hello and Welcome %s", *message)
	case "help":
		fmt.Println("Print some help message")
	default:
		fmt.Println("Default")
	}
}