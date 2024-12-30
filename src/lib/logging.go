package lib

import (
	"fmt"

	"uilliam.com/sip/types"
)

func GenerateLoggingCode(logging types.LoggingDefinition, path string) string {
	if !logging.Enabled {
		// If logging not enabled, return nothing
		return ""
	}

	codeSnippet := fmt.Sprintf(`
// Logging enabled with type=%s, level=%s
// (Pseudo-implementation)
log.Printf("Request to %s started\n")
`, logging.Type, logging.Level, path)

	return codeSnippet
}

// GenerateLoggingEnd returns a snippet for "logging end".
func GenerateLoggingEnd(logging types.LoggingDefinition, path string) string {
	if !logging.Enabled {
		return ""
	}
	return fmt.Sprintf(`
// Log request end
log.Printf("Request to %s ended")
`, path)
}
