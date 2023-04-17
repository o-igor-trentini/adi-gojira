package encoder

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// DecodeRequestResponse recebe a resposta de uma requisição e converte o body em struct.
func DecodeRequestResponse[T any](res *http.Response, dst *T) ([]byte, error) {
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("não foi possível ler o corpo da resposta [erro: %w]", err)
	}
	defer res.Body.Close()

	if err := json.Unmarshal(body, &dst); err != nil {
		return nil, fmt.Errorf("não foi possível converter os bytes para objeto [type: %T, erro: %w]", dst, err)
	}

	return body, nil
}

// Decode converte os bytes em struct.
func Decode[T any](body []byte, dst *T) error {
	if err := json.Unmarshal(body, &dst); err != nil {
		return fmt.Errorf("não foi possível converter os bytes para objeto [type: %T, erro: %w]", dst, err)
	}

	return nil
}
