package main

import (
	"fetch-receipt-processor/routes"

	"github.com/gin-gonic/gin"

	"log"
	"os"
	"net/http"
	"context"
	// "time"
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

	// config the server on 8080
	router := gin.Default()
	routes.ReceiptRoutes(router)
	var portNum = ":8080"
	server := &http.Server {
		Addr: portNum,
		Handler: router,
	}
	log.Printf("Attempt to start the server on %s.\n", portNum)

	// start the server
	go server.ListenAndServe()
	log.Println("Server start.")

	server.Shutdown(context.Background())
	log.Println("Server exiting")
}