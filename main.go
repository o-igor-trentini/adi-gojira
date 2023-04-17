package main

import (
	"fmt"
	"github.com/o-igor-trentini/adi-gojira/pkg/gjservice"
	"net/http"
	"strconv"
	"time"
)

func main() {
	httpClient := &http.Client{Timeout: time.Duration(10) * time.Second}

	developerCustomField := "Desenvolvedor"

	config := gjservice.Config{
		BaseURL:      "",
		JiraUsername: "",
		JiraToken:    "",
		HTTPClient:   *httpClient,
		CustomFields: gjservice.ConfigCustomFields{
			Developer: &developerCustomField,
		},
	}

	service := *gjservice.NewClient(config)

	params := map[string]string{
		"jql":        "project IN (PEC, RISK1) AND resolutiondate >= 2023-03-01 AND resolutiondate <= 2023-03-30",
		"maxResults": strconv.Itoa(100),
		"startAt":    "150",
		//"fields":     strings.Join(fields, ","),
		"expand": "names",
	}

	response, err := service.SearchByJQL(params)
	fmt.Println(response, err)
}
