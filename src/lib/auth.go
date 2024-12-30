package lib

import (
	"fmt"

	"uilliam.com/sip/types"
)

// GenerateAuthCode returns the snippet of Go code that implements auth logic
// loggingCode := lib.GenerateLoggingCode(cfg.Logging, cfg.Path)

func GenerateAuthCode(auth types.AuthDefinition) string {
	if auth.Type == "" {
		return ""
	}
	// For example:
	return fmt.Sprintf(`
	// Auth enabled: type=%s, roles=%v
	// (Pseudo-implementation)
	// if !isAuthorized(r, %v) {
	//     http.Error(w, "Unauthorized", http.StatusUnauthorized)
	//     return
	// }
	`, auth.Type, auth.Roles, auth.Roles)
}
