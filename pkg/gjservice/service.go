package gjservice

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/o-igor-trentini/adi-gojira/internal/encoder"
)

// SearchByJQL busca de forma paginada as tarefas de acordo com JQL (SQL do Jira).
func (c Client) SearchByJQL(queryParams map[string]string) (SearchByJQLPayload, error) {
	var qParams string

	if len(queryParams) > 0 {
		qParams += "?"

		for k, v := range queryParams {
			qParams += k + "=" + url.QueryEscape(v) + "&"
		}

		qParams = qParams[:len(qParams)-1]
	}

	data := SearchByJQLPayload{}

	_, body, err := c.get("search" + qParams)
	if err != nil {
		return data, fmt.Errorf("não foi possível buscar por JQL [erro: %s]", err)
	}

	var names expandedNames

	if err := encoder.Decode(body, &names); err != nil {
		return data, err
	}

	var developerFieldKeys []string
	devField := c.customFields.Developer

	// pega os campos customizados de 'desenvolvedor' de cada projeto
	// e transforma em uma única chave
	if len(names.Values) > 0 {
		if devField != nil {
			for key, value := range names.Values {
				if value == *devField {
					developerFieldKeys = append(developerFieldKeys, key)
				}
			}
		}
		strBody := string(body)

		for _, v := range developerFieldKeys {
			strBody = strings.ReplaceAll(string(body), v, "developers")
		}

		body = []byte(strBody)
	}

	return data, encoder.Decode(body, &data)
}
