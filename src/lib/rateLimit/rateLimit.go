package rateLimit

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"uilliam.com/sip/types"
)

// The `GenerateRateLimitCode` function generates rate limiting code based on the provided
// configuration and path.
//
// Author - Liam Scott
// Last update - 01/04/2025
//
// @ param config ()  - The `config` parameter in the `GenerateRateLimitCode` function is of type
// `types.RateLimitDefinition` and contains the configuration settings for rate limiting. It likely
// includes fields such as `Enabled` (a boolean indicating if rate limiting is enabled), `TimeWindow`
// (an integer representing
// @ param path (string)  - The `path` parameter in the `GenerateRateLimitCode` function represents the
// directory path where the rate-limiting code will be generated. This function generates rate-limiting
// code based on the provided configuration and saves it in the specified directory path.
func GenerateRateLimitCode(config types.RateLimitDefinition, path string) (string, []string) {

	if !config.Enabled {
		snippet := `func checkRateLimit(clientID string) bool {
	 return true 
	 }`
		return snippet, nil
	}

	rateLimitFileContent := `package ratelimit

import (
    "sync"
    "time"
)

// We'll track requests in memory for now, this will be changed in the future to use our connection module
// This map holds the timestamps of requests for each client (by IP or key).
// For a production scenario, you might use a more robust data structure or external store.
var requestsMap = make(map[string][]time.Time)
var mu sync.Mutex

// IsAllowed checks if a given client is allowed to make a request at this moment.
func IsAllowed(clientID string) bool {

    now := time.Now()
    cutoff := now.Add(-time.Duration(` + strconv.Itoa(config.TimeWindow) + `) * time.Second)

    mu.Lock()
    defer mu.Unlock()

    // Prune old timestamps
    validReqs := []time.Time{}
    for _, t := range requestsMap[clientID] {
        if t.After(cutoff) {
            validReqs = append(validReqs, t)
        }
    }
    requestsMap[clientID] = validReqs

    // Check if we're under the limit
    if len(validReqs) < ` + strconv.Itoa(config.Limit) + ` {
        // Allowed! Record the new request
        requestsMap[clientID] = append(requestsMap[clientID], now)
        return true
    }

    // Deny
    return false
}
`

	// 2. Ensure folder `ratelimit` exists at `path`.
	rateLimitFolder := filepath.Join(path, "ratelimit")
	if err := os.MkdirAll(rateLimitFolder, 0755); err != nil {
		return fmt.Sprintf("error : failed to create ratelimit folder: %v", err), nil
	}

	// 3. Create the file (ratelimit.go) and write the code
	rateLimitFilePath := filepath.Join(rateLimitFolder, "ratelimit.go")
	f, err := os.Create(rateLimitFilePath)
	if err != nil {
		return fmt.Sprintf("error : failed to create ratelimit.go: %v", err), nil
	}
	defer f.Close()

	if _, err := f.WriteString(rateLimitFileContent); err != nil {
		return fmt.Sprintf("error : failed to write ratelimit code: %v", err), nil
	}

	// 4. Generate a snippet that references this code
	// For instance, a function you can call in main.go that checks if a request is allowed

	snippet := `func checkRateLimit(clientID string) bool {
    return ratelimit.IsAllowed(clientID)
}`

	imports := []string{
		"uilliam.com/sip/ratelimit",
	}

	return snippet, imports
}
