package test

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"

	"uilliam.com/sip/lib/rateLimit"
	"uilliam.com/sip/types"
)

// Example usage:
// {
//   "ratelimit": {
//     "enabled": true,
//     "limit": 10,
//     "timeWindow": 60
//   }
// }

// The TestRateLimit function defines a rate-limit configuration, loads environment variables,
// generates rate-limiting code, and outputs the snippet and required imports.
//
// Author - Liam Scott
// Last update - 01/04/2025
func TestRateLimit() {
	// Define our rate-limit configuration
	config := types.RateLimitDefinition{
		Enabled:    true, // or false
		Limit:      10,   // max requests in the time window
		TimeWindow: 60,   // in seconds
	}

	// Load the .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		return
	}

	path := os.Getenv("build_path")

	// Call the GenerateRateLimitCode function
	rateLimitSnippet, imports := rateLimit.GenerateRateLimitCode(config, path)

	// If there's an error, check for the prefix "error :"
	if len(rateLimitSnippet) >= 7 && rateLimitSnippet[:7] == "error :" {
		log.Fatalf("Error: %v", rateLimitSnippet)
	}

	// Print out the snippet and imports. In a real application,
	// you would integrate them into your main.go as needed.
	fmt.Println("========== Generated Snippet ==========")
	fmt.Println(rateLimitSnippet)
	fmt.Println("========== Required Imports ==========")
	for _, imp := range imports {
		fmt.Println(imp)
	}

	fmt.Println()
	fmt.Printf("RateLimit code has been generated under '%s/ratelimit/ratelimit.go'\n", path)
	fmt.Println("Include the snippet in your main.go and the imports to use it!")
}
