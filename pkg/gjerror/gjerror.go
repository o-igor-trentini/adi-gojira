package gjerror

import (
	"fmt"
	"strings"
)

// APIError é formato da resposta quando ocorre algum erro na requisição.
type APIError interface {
	Error() string
	GetCode() Code
	GetMessage() Code
	GetErrorMessages() []string
	GetWarningMessages() []string
}

type APIErrorImpl struct {
	Code            Code
	Message         Message
	ErrorMessages   []string `json:"errorMessages"`
	WarningMessages []string `json:"warningMessages"`
}

func (e APIErrorImpl) Error() string {
	msg := fmt.Sprintf("[%s] %s", e.Code, e.Message)

	if errs := e.ErrorMessages; len(errs) > 0 {
		msg += fmt.Sprintf("\n\nErros: %s", strings.Join(errs, ", "))
	}

	if warnings := e.WarningMessages; len(warnings) > 0 {
		msg += fmt.Sprintf("\n\nAvisos: %s", strings.Join(warnings, ", "))
	}

	return msg
}

func (e APIErrorImpl) GetCode() Message {
	return e.Message
}

func (e APIErrorImpl) GetMessage() Code {
	return e.Code
}

func (e APIErrorImpl) GetErrorMessages() []string {
	return e.ErrorMessages
}

func (e APIErrorImpl) GetWarningMessages() []string {
	return e.WarningMessages
}
