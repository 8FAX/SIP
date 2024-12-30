package lib

import (
	"fmt"
	"strings"

	"uilliam.com/sip/types"
)

// GeneratePluginsCode generates the snippet for a given plugin trigger
func GeneratePluginsCode(plugins []types.PluginDefinition, trigger string) string {
	var sb strings.Builder
	var found bool

	for _, p := range plugins {
		if p.Trigger == trigger {
			found = true
			sb.WriteString(fmt.Sprintf(`    // Plugin: %s (%s)
    // (Pseudo-execution)
`, p.Name, p.Script))

			if strings.ToLower(p.Trigger) == "passthrough" && p.Data != "" {
				sb.WriteString(fmt.Sprintf(`    // data to pass: %s
`, p.Data))
			}
		}
	}

	if found {
		return fmt.Sprintf("\n    // Execute %s plugins\n%v", trigger, sb.String())
	}
	return ""
}
