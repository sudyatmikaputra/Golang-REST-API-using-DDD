package internal

type ParameterType string

const (
	All       ParameterType = "all"
	Doctor    ParameterType = "doctor"
	Merchant  ParameterType = "merchant"
	Medicplus ParameterType = "medicplus"
)

// errors
const (
	ErrInvalidRequest = "invalid request"
)
