package gjservice

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/o-igor-trentini/adi-gojira/internal/encoder"
)

// SearchByJQL busca de forma paginada as tarefas de acordo com JQL (SQL do Jira).
func (c Client) SearchByJQL(queryParams map[string]string, customFields CustomFields) (
	SearchByJQLPayload,
	error,
) {
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

	// pega os campos customizados de 'desenvolvedor' de cada projeto
	// e transforma em uma única chave
	if len(customFields.Developer) > 0 {
		strBody := string(body)

		for _, v := range customFields.Developer {
			strBody = strings.ReplaceAll(strBody, v, "developers")
		}

		body = []byte(strBody)
	}

	return data, encoder.Decode(body, &data)
}
