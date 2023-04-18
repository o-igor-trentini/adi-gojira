package encoder

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Decode converte os bytes em struct.
func Decode[T any](body []byte, dst *T) error {
	if err := json.Unmarshal(body, &dst); err != nil {
		return fmt.Errorf("não foi possível converter os bytes para objeto [type: %T, erro: %w]", dst, err)
	}

	return nil
}

// ResponseBodyToBytes converte o corpo a resposta da requisição para um array de bytes e fecha o res.Body.
func ResponseBodyToBytes(res *http.Response) ([]byte, error) {
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("não foi possível ler o corpo da resposta [erro: %w]", err)
	}
	defer res.Body.Close()

	return body, err
}

// DecodeRequestResponse recebe a resposta de uma requisição e converte o body em struct.
func DecodeRequestResponse[T any](res *http.Response, dst *T) ([]byte, error) {
	body, err := ResponseBodyToBytes(res)
	if err != nil {
		return nil, err
	}

	if err := Decode(body, &dst); err != nil {
		return nil, err
	}

	return body, nil
}
