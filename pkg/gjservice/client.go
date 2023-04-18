package gjservice

import (
	"fmt"
	"github.com/o-igor-trentini/adi-gojira/internal/encoder"
	"github.com/o-igor-trentini/adi-gojira/pkg/gjerror"
	"net/http"

	"github.com/o-igor-trentini/adi-gojira/internal/gjutils"
)

type Client struct {
	baseURL    string
	baseAuth   string
	httpClient http.Client
}

// NewClient instância o cliente que realiza as requisições para a API do Jira Software.
func NewClient(config Config) *Client {
	return &Client{
		baseURL:    config.BaseURL + "/rest/api/3/",
		baseAuth:   gjutils.BasicAuth(config.JiraUsername, config.JiraToken),
		httpClient: config.HTTPClient,
	}
}

func (c Client) addAuthorizationHeader(req *http.Request) {
	req.Header.Add("Authorization", "Basic "+c.baseAuth)
}

func (c Client) get(path string) (*http.Response, []byte, error) {
	url := c.baseURL + path

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, fmt.Errorf("não foi possível criar a requisição '%s' [erro: %w]", url, err)
	}

	c.addAuthorizationHeader(req)

	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, nil, fmt.Errorf("não foi possível completar requisição para '%s' [erro: %w]", url, err)
	}

	body, err := c.handleResponseError(res)

	return res, body, err
}

func (c Client) handleResponseError(res *http.Response) ([]byte, error) {
	var responseErr gjerror.APIErrorImpl

	if res.StatusCode >= 400 {
		switch res.StatusCode {
		case http.StatusTooManyRequests:
			responseErr.Code = gjerror.CodeRateLimit

			detail := "alguns"
			if v, ok := res.Header["Retry-After"]; ok && len(v) > 0 {
				detail = v[0]
			}

			responseErr.Message = gjerror.Message(fmt.Sprintf(string(gjerror.MessageRateLimit), detail))

		default:
			body, err := encoder.DecodeRequestResponse(res, &responseErr)
			if err != nil {
				return body, fmt.Errorf("não foi possível validar a resposta da requisição [erro: %w]", err)
			}

			responseErr.Code = gjerror.CodeDefault
			responseErr.Message = gjerror.MessageDefault
		}
	}

	return encoder.ResponseBodyToBytes(res)
}
