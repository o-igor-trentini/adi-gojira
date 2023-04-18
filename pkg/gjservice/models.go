package gjservice

import (
	"github.com/o-igor-trentini/adi-gojira/pkg/gjmodels"
)

type SearchByJQLPayload struct {
	gjmodels.Pagination
	Issues []gjmodels.Issue `json:"issues"`
}

type CustomFields struct {
	Developer []string
}
