package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/manzurahammed/rm-cli/client"
)
var (
	backendAPIURL = flag.String("backend","http://localhost:9000","Bacakend Url")
	helpFlag = flag.Bool("help",false,"Help Flag")
)

func main(){
	flag.Parse()
	s := client.NewSwitch(*backendAPIURL)
	if * helpFlag || len(os.Args)==1 {
		s.Help()
		return 
	}
	err := s.Switch()

	if err!=nil {
		fmt.Println("Error found")
		os.Exit(2)
	}
}