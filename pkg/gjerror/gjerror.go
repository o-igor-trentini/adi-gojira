package gjerror

import (
	"fmt"
	"strings"
)

// APIError é formato da resposta quando ocorre algum erro na requisição.
type APIError interface {
	Error() string
	GetErrorMessages() []string
	GetWarningMessages() []string
}

type APIErrorImpl struct {
	ErrorMessages   []string `json:"errorMessages"`
	WarningMessages []string `json:"warningMessages"`
}

func (e APIErrorImpl) Error() string {
	msg := fmt.Sprintf("Erros: %s", strings.Join(e.GetErrorMessages(), ", "))

	if warnings := e.GetErrorMessages(); len(warnings) > 0 {
		msg += fmt.Sprintf(" || Avisos: %s", strings.Join(warnings, ", "))
	}

	return msg
}

func (e APIErrorImpl) GetErrorMessages() []string {
	return e.ErrorMessages
}

func (e APIErrorImpl) GetWarningMessages() []string {
	return e.WarningMessages
}
