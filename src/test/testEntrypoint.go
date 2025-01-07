package test

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"

	"uilliam.com/sip/lib/entrypoint"
	"uilliam.com/sip/types"
)

/*
TestEntrypoint demonstrates generating a main.go file using
our GenerateMainEntrypoint function. It creates a small set of
endpoint configs, then writes the final output to your build_path.

Author - Liam Scott
Last update - 01/04/2025
*/
func TestEntrypoint() {
	// Load the .env file (e.g., to get build_path)
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: no .env file found or error loading it.")
	}

	// Where to place the generated code (e.g. "./build")
	buildPath := os.Getenv("build_path")
	if buildPath == "" {
		buildPath = "./build"
	}

	// Create a sample ConfigDoc (top-level metadata)
	config := types.ConfigDoc{
		Name:    "MyAPI",
		Version: "v1.0",
		Threads: 3, // number of global threads (if you need it)
	}

	// Define endpoints (keys can be anything, e.g. "getUser", "getinfo", etc.)
	endpoints := map[string]types.EndpointConfig{
		"getUser": {
			Path:    "/getUser",
			Method:  "GET",
			Threads: 2,
		},
		"getinfo": {
			Path:    "/getinfo",
			Method:  "GET",
			Threads: 4,
		},
	}

	// Generate the main.go entrypoint
	err = entrypoint.GenerateMainEntrypoint(config, endpoints, buildPath)
	if err != nil {
		log.Fatalf("Error generating main.go: %v", err)
	}

	fmt.Println("Main entrypoint generation complete!")
	fmt.Printf("Check the file at: %s/main.go\n", buildPath)
}
