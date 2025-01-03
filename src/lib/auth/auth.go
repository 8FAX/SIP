package auth

import (
	auth "uilliam.com/sip/lib/auth/gen"
	"uilliam.com/sip/types"
)

// GenerateAuthCode returns the snippet of Go code that implements auth logic

// "type": "APIKey", // none, APIKey, OAuth2

func GenerateAuthCode(authType types.AuthDefinition) (string, []string) {
	if authType.Type == "none" {
		return " // No auth required\n", nil
	}

	if authType.Type == "APIKey" {
		return auth.GenerateAPIKeyCode(authType)
	}

	if authType.Type == "OAuth2" {
		return auth.GenerateOAuth2Code(authType)
	}

	return " // No auth required\n", nil

}
