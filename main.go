package main

import (
	"os"
	"log"
)

func main() {
	// create the log file with read and write permission
	file, err := os.OpenFile("dev.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	log.SetOutput(file)
	log.Println("Set log output to dev.log file")

	
}