package types

// This file holds the definitions of your structs so that they
// can be imported by both main.go and the lib packages.

type ParamDefinition struct {
	Type     string `json:"type"`
	Required bool   `json:"required"`
}

type RequestDefinition struct {
	Params  map[string]ParamDefinition `json:"params"`
	Headers map[string]ParamDefinition `json:"headers"`
}

type ResponseDefinition struct {
	Fields        map[string]string `json:"fields"`
	ErrorHandling map[string]string `json:"errorHandling"`
}

type QueryDefinition struct {
	Type      string `json:"type"`
	Statement string `json:"statement"`
}

type DatabaseDefinition struct {
	Connection string          `json:"connection"`
	Operation  string          `json:"operation"`
	Query      QueryDefinition `json:"query"`
}

type PluginDefinition struct {
	Name    string `json:"name"`
	Trigger string `json:"trigger"`
	Data    string `json:"data,omitempty"`
	Script  string `json:"script"`
}

type AuthDefinition struct {
	Type        string           `json:"type"`
	Roles       []string         `json:"roles"`
	KeyLocation string           `json:"keyLocation"`
	Parameters  AuthAPIKeyConfig `json:"parameters"`
}

type AuthAPIKeyConfig struct {
	KeyName    string                 `json:"keyName"`
	Connection APIKeyConfigConnection `json:"connection"`
}

type APIKeyConfigConnection struct {
	Name  string `json:"name"`
	Query string `json:"query"`
}

type LoggingDefinition struct {
	Enabled    bool   `json:"enabled"`
	Type       string `json:"type"`
	Connection string `json:"connection"`
	Level      string `json:"level"`
}

type LoggingLevels struct {
	Info  int
	Debug int
	Error int
	Warn  int
}

type CacheDefinition struct {
	Enabled    bool   `json:"enabled"`
	Type       string `json:"type"`
	Connection string `json:"connection"`
	TTL        int    `json:"ttl"`
}

type RateLimitDefinition struct {
	Limit      int `json:"limit"`
	TimeWindow int `json:"timeWindow"`
}

type EndpointConfig struct {
	Path        string              `json:"path"`
	Method      string              `json:"method"`
	Description string              `json:"description"`
	Request     RequestDefinition   `json:"request"`
	Response    ResponseDefinition  `json:"response"`
	Database    DatabaseDefinition  `json:"database"`
	Plugins     []PluginDefinition  `json:"plugins"`
	Auth        AuthDefinition      `json:"auth"`
	Logging     LoggingDefinition   `json:"logging"`
	Cache       CacheDefinition     `json:"cache"`
	RateLimit   RateLimitDefinition `json:"rateLimit"`
	Tags        []string            `json:"tags"`
	Version     string              `json:"version"`
}
