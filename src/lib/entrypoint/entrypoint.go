package entrypoint

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"uilliam.com/sip/types"
)

// The function `GenerateMainEntrypoint` generates the main.go file with the main code for a Go
// application based on the provided configuration and endpoints.
//
// Author - Liam Scott
// Last update - 01/06/2025
//
// @ param config ()  - The `config` parameter is of type `types.ConfigDoc` and contains configuration
// information for the main entry point generation function. It likely includes settings and
// configurations needed for the application to run.
// @ param endpoints ()  - The `endpoints` parameter in the `GenerateMainEntrypoint` function is a map
// that contains information about different endpoints in your application. Each key in the map
// represents the endpoint's name or identifier, and the corresponding value is an `EndpointConfig`
// struct that holds configuration details for that endpoint
// @ param buildPath (string)  - The `buildPath` parameter is a string that represents the directory
// path where the main.go file will be generated.
//
// @ returns The `GenerateMainEntrypoint` function returns an error.
func GenerateMainEntrypoint(config types.ConfigDoc, endpoints map[string]types.EndpointConfig, buildPath string) error {
	mainCode := buildMainGo(config, endpoints)

	mainFilePath := filepath.Join(buildPath, "main.go")
	if err := os.WriteFile(mainFilePath, []byte(mainCode), 0644); err != nil {
		return fmt.Errorf("failed to write main.go: %w", err)
	}

	return nil
}

// The `buildMainGo` function generates a main Go file with imports, struct definitions, main function,
// endpoint channels, thread handling, HTTP server setup, and request routing for multiple endpoints.
//
// Author - Liam Scott
// Last update - 01/06/2025
//
// @ param config ()  - The `config` parameter in the `buildMainGo` function represents a configuration
// document containing information about the application. This information typically includes the name
// and version of the application. The `types.ConfigDoc` type likely defines a structure to hold this
// configuration data.
// @ param endpoints ()  - {
//
// @ returns The function `buildMainGo` returns a string containing a generated Go code snippet for a
// main package. The generated code includes imports, struct definitions, main function setup,
// initialization of endpoint channels, starting threads for each endpoint, setting up HTTP server for
// dispatch, routing requests to specific endpoints, and starting the server to listen on port 8080.
func buildMainGo(config types.ConfigDoc, endpoints map[string]types.EndpointConfig) string {
	sb := &strings.Builder{}

	// Basic imports
	sb.WriteString(`package main

import (
    "fmt"
    "log"
    "net/http"
    "time"
)
`)

	// Add endpoint imports
	for name := range endpoints {
		pkgName := sanitizePkgName(name)
		sb.WriteString(fmt.Sprintf("import \"%s/%s\"\n", "uilliam.com/sip", pkgName))
	}

	// Define a struct to store each endpoint’s channel
	sb.WriteString(`
type EndpointChannels struct {
`)
	for name := range endpoints {
		pkgName := sanitizePkgName(name)
		sb.WriteString(fmt.Sprintf("    %sChan chan %s.Request\n", pkgName, pkgName))
	}
	sb.WriteString("}\n\n")

	// main function
	sb.WriteString("func main() {\n")

	sb.WriteString(fmt.Sprintf("    fmt.Println(\"Starting %s version %s...\")\n", config.Name, config.Version))

	// Initialize EndpointChannels
	sb.WriteString("    channels := &EndpointChannels{\n")
	for name := range endpoints {
		pkgName := sanitizePkgName(name)
		sb.WriteString(fmt.Sprintf("        %sChan: make(chan %s.Request, 100),\n", pkgName, pkgName))
	}
	sb.WriteString("    }\n\n")

	// Start threads for each endpoint
	for name, spec := range endpoints {
		pkgName := sanitizePkgName(name)
		sb.WriteString(fmt.Sprintf("    // Spin up %d threads for endpoint %q\n", spec.Threads, name))
		sb.WriteString(fmt.Sprintf("    for i := 0; i < %d; i++ {\n", spec.Threads))
		sb.WriteString(fmt.Sprintf("        go %s.HandleRequest(channels.%sChan)\n", pkgName, pkgName))
		sb.WriteString("    }\n\n")
	}

	// Set up HTTP (or any other) server for dispatch
	sb.WriteString(`
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        path := r.URL.Path
        switch path {
`)

	// For each endpoint, route the request
	for name, spec := range endpoints {
		pkgName := sanitizePkgName(name)
		sb.WriteString(fmt.Sprintf("        case %q:\n", spec.Path))
		sb.WriteString("            // Construct a Request object\n")
		sb.WriteString(fmt.Sprintf("            req := %s.Request{\n", pkgName))
		sb.WriteString("                ID: r.URL.Query().Get(\"uuid\"),\n")
		sb.WriteString("                Time: time.Now(),\n")
		sb.WriteString("            }\n")
		sb.WriteString(fmt.Sprintf("            channels.%sChan <- req\n", pkgName))
		sb.WriteString(fmt.Sprintf("            fmt.Fprintf(w, \"Enqueued request for %s\")\n", name))
		sb.WriteString("            return\n")
	}

	// Default not found
	sb.WriteString(`
        default:
            http.NotFound(w, r)
        }
    })
`)

	sb.WriteString(fmt.Sprintf("    log.Fatal(http.ListenAndServe(\":%d\", nil))\n", 8080)) // Hardcoded or from config
	sb.WriteString("}\n")

	return sb.String()
}

// The `sanitizePkgName` function replaces all occurrences of "-" with "_" in a given string.
//
// Author - Liam Scott
// Last update - 01/06/2025
//
// @ param name (string)  - The `sanitizePkgName` function takes a string `name` as input and replaces
// all occurrences of "-" with "_".
//
// @ returns The `sanitizePkgName` function is returning a modified version of the input `name` string
// where any occurrences of "-" are replaced with "_".
func sanitizePkgName(name string) string {
	return strings.ReplaceAll(name, "-", "_")
}
