package logging

import (
	"fmt"
	"os"
	"path/filepath"

	"uilliam.com/sip/types"
)

// GenerateLoggingCode generates the Go code for logging based on the provided logging definition.

// Thoughts:
// - unbuffered design so if logs spike it won't block the main thread.
// - we dont need to open the file every time, we can keep it open and write to it.
// add a graceful shutdown to close the file when the program exits.
// concurrency may become an issue if we have multiple goroutines writing to the same file.
// we may end up needing a mutex to lock the file while writing to it or
// we can use a second buffer to stor all the logs and flush them to the file at a set interval or
// we have a dedicated goroutine that writes to the file and the other goroutines send messages to it.
// i dont want to overcomplicate it so we will keep it simple for now, with the file being opened and closed every time we write to it
// this is not foolproof but it will work for now, as we can still run into issues if we have multiple goroutines writing to the file at the same time.
// we can always come back and improve it later.

// "logging": {
//       "enabled": true,
//       "type": "File", // or "Database" or "Console"
//         "connection": "MyLog", // can be a connection name or none
//       "level": "info"
//     },

// The `GenerateLoggingCode` function generates logging code based on a provided logging definition and
// saves it to a specified path, returning a snippet for `main.go` and necessary imports.
//
// Author - Liam Scott
// Last update - 01/03/2025
//
// @ param logging ()  - - Level: "debug"
// @ param path (string)  - The `path` parameter in the `GenerateLoggingCode` function represents the
// directory path where the logging code will be generated. This function generates logging code based
// on the provided `logging` configuration and saves it in the specified directory. The logging code
// includes functionality to log messages to different destinations based on the
func GenerateLoggingCode(logging types.LoggingDefinition, path string) (string, []string) {

	// 1) Build the logger worker code
	loggerWorkerCode := `package logger

import (
	"fmt"
	"os"
	"strings"
	"time"
)

// LoggingLevels maps log-level strings to numeric priorities.
var LoggingLevels = map[string]int{
    "error": 1,
    "warn":  2,
    "info":  3,
    "debug": 4,
}

// ConfiguredLevel is the globally configured logging level.
var ConfiguredLevel = "` + logging.Level + `"

// LogChannel is the global channel for passing log messages to the logger worker.
var LogChannel chan string

// InitLogger initializes the logging channel and starts 'threads' logger workers.
func InitLogger(threads int) {
	LogChannel = make(chan string, 1000)
	for i := 0; i < threads; i++ {
		go loggerWorker()
	}
}

// loggerWorker runs in a loop, processing log messages from LogChannel.
func loggerWorker() {
	for {
		select {
		case logMsg := <-LogChannel:
			// Process the log message
			handleLog(logMsg)
		}
	}
}

// handleLog processes a log message based on the configured logging definition.
func handleLog(logMsg string) {
	// Convert the configured level string to a numeric
	configVal, ok := LoggingLevels[ConfiguredLevel]
	if !ok {
		// If the configured level is unknown, default to 'info'
		configVal = LoggingLevels["info"]
	}

	var messageLevel string
	var actualMessage string

	// Simple parse: find the first colon
	colonIndex := strings.Index(logMsg, ":")
	if colonIndex <= 0 {
		// If there's no level prefix, default to info
		messageLevel = "info"
		actualMessage = logMsg
	} else {
		messageLevel = strings.TrimSpace(logMsg[:colonIndex])
		actualMessage = strings.TrimSpace(logMsg[colonIndex+1:])
	}

	// Convert messageLevel to numeric priority (default to "info" if invalid)
	msgVal, ok := LoggingLevels[messageLevel]
	if !ok {
		msgVal = LoggingLevels["info"]
	}

	// Only log if the message’s level <= configured level
	if msgVal <= configVal {
		// Format the final string with timestamp + alignment
		currentTime := time.Now().Format("2006-01-02 15:04:05")
		alignTime := 26 - len(currentTime)
		if alignTime < 0 {
			alignTime = 0
		}

		alignLevel := 10 - len(messageLevel)
		if alignLevel < 0 {
			alignLevel = 0
		}

		// Build the aligned log message
		formattedLog := fmt.Sprintf("[%s%s]:[%s%s]: %s",
			currentTime, strings.Repeat(" ", alignTime),
			messageLevel, strings.Repeat(" ", alignLevel),
			actualMessage,
		)

		// <LOGGING_TYPE_SWITCH>
		// This will be replaced or updated by code generation:
`

	// Now append the switch logic:
	switch logging.Type {
	case "File":
		loggerWorkerCode += `
		// Logging to a file
		f, err := os.OpenFile("` + logging.Connection + `", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Println("Error opening log file:", err)
			return
		}
		defer f.Close()

		if _, err := f.WriteString(formattedLog + "\n"); err != nil {
			fmt.Println("Error writing to log file:", err)
		}
`
	case "Console":
		loggerWorkerCode += `
		// Logging to console
		fmt.Println(formattedLog)
`
	case "connection":
		loggerWorkerCode += `
		// Logging to a network/db/connection (not implemented)
		fmt.Println("Log to connection not implemented:", formattedLog)
`
	default:
		loggerWorkerCode += `
		// Default to console
		fmt.Println(formattedLog)
`
	}

	// Close out the handleLog function
	loggerWorkerCode += `
	} // End of "if msgVal <= configVal"
}
`

	// 2) Write the logger worker code to <path>/logger/logger.go

	// Ensure the "logger" subfolder exists
	loggerFolderPath := filepath.Join(path, "logger")
	if err := os.MkdirAll(loggerFolderPath, 0755); err != nil {
		return "error :" + fmt.Sprintf("failed to create logger folder: %w", err), nil
	}

	// Create (or overwrite) the logger.go file
	loggerFilePath := filepath.Join(loggerFolderPath, "logger.go")
	f, err := os.Create(loggerFilePath)
	if err != nil {
		return "error :" + fmt.Sprintf("failed to create logger.go: %w", err), nil
	}
	defer f.Close()

	// Write the worker code to logger.go
	if _, err := f.WriteString(loggerWorkerCode); err != nil {
		return "error :" + fmt.Sprintf("failed to write logger code: %w", err), nil
	}

	// 3) Generate the snippet for main.go that we return as a string.
	//    This snippet references the global `logger.LogChannel` from logger/logger.go.
	//    We also embed the "level" as a prefix so the worker knows the message’s level.
	mainLoggerSnippet := `func logger(level, message string) {
    // If disabled, skip
    if !` + fmt.Sprintf("%v", logging.Enabled) + ` {
        return
    }
    // Example: embed the level as a prefix so loggerWorker can parse it
    logMsg := fmt.Sprintf("%s: %s", level, message)
    logger.LogChannel <- logMsg
}`
	// The imports needed for this snippet
	imports := []string{
		"fmt",
		"uilliam.com/sip/logger",
	}

	return mainLoggerSnippet, imports
}
