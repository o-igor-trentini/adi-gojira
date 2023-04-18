package gjerror

type Code string

type Message string

const (
	CodeRateLimit    Code    = "rate-limit"
	MessageRateLimit Message = "Limite de requisições atingido, tente novamente em %s segundos"

	CodeDefault    Code    = "other"
	MessageDefault Message = "Não foi possível realizar a requisição"
)
