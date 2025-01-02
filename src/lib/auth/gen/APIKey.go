package auth

import (
	"uilliam.com/sip/types"
)

// GenerateAPIKeyCode returns the snippet of Go code that implements APIKey logic

// "auth": {
// 	"type": "APIKey", // none, APIKey, OAuth2
// 	"keylocation": "header", // or "query" or "body"
// 	"peramaeters": {
// 	  "keyName": "X-API-Key",
// 	  "connection": {
// 		"name": "MyDatabase",
// 		"query": "SELECT apiKey FROM users WHERE id = :uuid"
// 	  }
// 	},

// The function `GenerateAPIKeyCode` generates a code snippet in Go for validating an API key based on
// the provided authentication parameters.
//
// Author - Liam Scott
// Last update - 01/02/2025
//
// @ param auth ()  - The `GenerateAPIKeyCode` function takes an `auth` parameter of type
// `types.AuthDefinition`, which contains information about API key authentication configuration. The
// function generates a code snippet for validating API keys based on the provided configuration.
func GenerateAPIKeyCode(auth types.AuthDefinition) (string, []string) {
	var apiKeyConfig = types.AuthAPIKeyConfig{
		KeyName:    auth.Parameters.KeyName,
		Connection: auth.Parameters.Connection,
	}

	keyName := apiKeyConfig.KeyName

	code := "" // Initialize code snippet

	imports := []string{}

	// Add function header
	code += "func ValidateAPIKey(r *http.Request) (*http.Request, bool, string) {\n"

	// Add imports
	imports = append(imports, "net/http")
	imports = append(imports, "encoding/json")
	imports = append(imports, "uilliam.com/sip/connections")

	// Extract the key based on location
	code += "\tswitch \"" + auth.KeyLocation + "\" {\n"
	code += "\tcase \"header\":\n"
	code += "\t\tkey := r.Header.Get(\"" + keyName + "\")\n"
	code += "\tcase \"query\":\n"
	code += "\t\tkey := r.URL.Query().Get(\"" + keyName + "\")\n"
	code += "\tcase \"body\":\n"
	code += "\t\tvar bodyKey string\n"
	code += "\t\tif err := json.NewDecoder(r.Body).Decode(&bodyKey); err != nil {\n"
	code += "\t\t\treturn r, false, \"Invalid request body\"\n"
	code += "\t\t}\n"
	code += "\t\tkey := bodyKey\n"
	code += "\t}\n"

	// Check if key is empty
	code += "\tif key == \"\" {\n"
	code += "\t\treturn r, false, \"Missing API Key\"\n"
	code += "\t}\n"

	// Lookup key using connection library
	code += "\t// Use connection library to validate the key\n"
	code += "\tconnection := connections.GetConnection(\"" + apiKeyConfig.Connection.Name + "\")\n"
	code += "\tif connection == nil {\n"
	code += "\t\treturn r, false, \"Database connection not found\"\n"
	code += "\t}\n"
	code += "\tquery := \"" + apiKeyConfig.Connection.Query + "\"\n"
	code += "\tvalidKey, err := connection.ExecuteQuery(query, map[string]interface{}{})\n"
	code += "\tif err != nil {\n"
	code += "\t\treturn r, false, \"Error during key validation\"\n"
	code += "\t}\n"
	code += "\tif validKey != key {\n"
	code += "\t\treturn r, false, \"Invalid API Key\"\n"
	code += "\t}\n"

	// Return success
	code += "\treturn r, true, \"API Key verified successfully\"\n"
	code += "}\n"

	return code + "\n", imports
}
