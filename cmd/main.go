package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/VinnieApps/spa-server/pkg/handler"
)

func main() {
	server := http.NewServeMux()
	port := 80

	currentDirectory, err := os.Getwd()
	if err != nil {
		log.Fatal("Error while finding current directory.", err)
	}

	log.Printf("Serving files from '%s'\n", currentDirectory)
	server.HandleFunc("/", handler.MakeHandler(currentDirectory))

	log.Printf("Starting proxy at: http://localhost:%d\n", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), server); err != nil {
		log.Fatal(err)
	}
}
