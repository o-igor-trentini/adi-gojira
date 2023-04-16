package gjservice

import (
	"fmt"
	"net/url"

	"github.com/o-igor-trentini/adi-gojira/internal/encoder"
	"github.com/o-igor-trentini/adi-gojira/pkg/gjmodels"
)

type SearchByJQLPayload struct {
	gjmodels.Pagination
	Issues []gjmodels.Issue `json:"issues"`
}

// SearchByJQL busca de forma paginada as tarefas de acordo com o JQL (SQL do Jira).
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

	res, err := c.get("search" + qParams)
	if err != nil {
		return data, fmt.Errorf("não foi possível buscar por JQL [erro: %s]", err)
	}

	return data, encoder.DecodeRequestResponse(res, &data)
}
