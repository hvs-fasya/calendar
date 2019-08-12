package main

import (
	"log"

	"github.com/hvs-fasya/calendar/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
