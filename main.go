package main

import (
	"flag"
	"fmt"
	"log"
)

func main() {

	pathPtr := flag.String("path", "./secrets.csv", "path to secrets file")
	projectNamePtr := flag.String("project", "", "name of gcp project")
	createSecretsPtr := flag.Bool("create", false, "whether to create secrets")

	flag.Parse()

	// Check if the project name is provided
	if *projectNamePtr == "" {
		log.Fatal("Error: Project name must be specified")
	}

	// Read CSV file and check for errors
	records, err := readCsvFile(*pathPtr)
	if err != nil {
		log.Fatalf("Error reading CSV file: %v", err)
	}

	fmt.Println(records)

	fmt.Println(*projectNamePtr)

	//TODO: Add GCP auth

	if *createSecretsPtr {
		createGcpSecrets(records, *projectNamePtr)
	}

	loadGcpSecrets(records, *projectNamePtr)
}
