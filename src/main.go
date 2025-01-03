package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"

	"uilliam.com/sip/lib/logging"
	"uilliam.com/sip/types"
)

// "logging": {
//       "enabled": true,
//       "type": "File", // or "Database" or "Console"
//         "connection": "MyLog", // can be a connection name or none
//       "level": "info"
//     },

func main() {

	config := types.LoggingDefinition{
		Enabled:    true,
		Type:       "File",
		Level:      "info",
		Connection: "MyLog.log",
	}

	// Load the .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		return
	}

	path := os.Getenv("build_path")

	// Call the GenerateLoggingCode function
	loggerSnippet, imports := logging.GenerateLoggingCode(config, path)

	// If there's an error string, let's check. We used a quick hack that starts with "error :" if something failed.
	if len(loggerSnippet) >= 7 && loggerSnippet[:7] == "error :" {
		log.Fatalf("Error: %v", loggerSnippet)
	}

	// Print out the snippet and imports. In a real application,
	// you would integrate them into your main.go as needed.
	fmt.Println("========== Generated Snippet ==========")
	fmt.Println(loggerSnippet)
	fmt.Println("========== Required Imports ==========")
	for _, imp := range imports {
		fmt.Println(imp)
	}

	fmt.Println()
	fmt.Printf("Logger code has been generated under '%s/logger/logger.go'\n", path)
	fmt.Println("Include the snippet in your main.go and the imports to use it!")
}
