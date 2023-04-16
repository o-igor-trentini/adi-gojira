package gjutils

import (
	"encoding/base64"
)

// BasicAuth gera uma autenticação básica pelo usuário e senha.
func BasicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}
