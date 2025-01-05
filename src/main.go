package main

import (
	"fmt"

	"uilliam.com/sip/test"
)

// The main function initiates the code generation process by testing rate limiting and logging
// functionalities.
//
// Author - Liam Scott
// Last update - 01/04/2025
func main() {
	fmt.Println("Starting the code generation process...")
	fmt.Println("=======================================")
	test.TestRateLimit()
	fmt.Println("=======================================")
	test.TestLogging()
	fmt.Println("=======================================")
	fmt.Println("All test done!")

}
