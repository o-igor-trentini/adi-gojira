package gjservice

import "net/http"

type Config struct {
	BaseURL      string
	JiraUsername string
	JiraToken    string
	HTTPClient   http.Client
	CustomFields ConfigCustomFields
}

type ConfigCustomFields struct {
	Developer *string
}
