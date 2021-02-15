package main

import (
	"fmt"
	"os"
	"strings"
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
		msg := "Remineds CLI -- cli basics"
		if len(osParam) >2 {
			f := strings.Split(osParam[2],"=")
		
			if len(f)==2 && f[0] == "--msg" {
				msg = f[1]
			}
		}
		fmt.Println("hello and welcome: %s\n", msg)
	case "help":
		fmt.Println("Print some help message")
	default:
		fmt.Println("Default")
	}
}