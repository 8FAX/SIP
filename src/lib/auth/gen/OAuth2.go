package auth

import (
	"uilliam.com/sip/types"
)

// GenerateOAuth2Code returns the snippet of Go code that implements OAuth2 logic

// {
// 	"type": "OAuth2",
// 	"keyLocation": "header",
// 	"parameters": {
// 	  "headerName": "X-Custom-Auth",
// 	  "tokenPrefix": "Token",
// 	  "TokenIntrospectionURL": "https://example.com/oauth2/introspect",
// 	  "ClientID": "myClientID",
// 	  "ClientSecret": "myClientSecret",
// 	}
//   }

func GenerateOAuth2Code(auth types.AuthDefinition) (string, []string) {
	warning := "OAuth2 is not implemented yet"

	return warning, nil
}
