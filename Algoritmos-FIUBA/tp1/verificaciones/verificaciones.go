package verificaciones

import "strconv"

const (
	OP_SUMA           = "+"
	OP_RESTA          = "-"
	OP_MULTIPLICACION = "*"
	OP_DIVISION       = "/"
	OP_RAIZ_CUADRADA  = "sqrt"
	OP_EXPONENCIACION = "^"
	OP_LOGARITMO      = "log"
	OP_TERNARIO       = "?"
)

func VerificarToken(token string) bool {
	_, err := strconv.Atoi(token)
	if err == nil {
		return true
	}
	switch token {
	case OP_SUMA, OP_RESTA, OP_MULTIPLICACION, OP_DIVISION, OP_EXPONENCIACION, OP_RAIZ_CUADRADA, OP_LOGARITMO, OP_TERNARIO:
		return true
	default:
		return false
	}
}

func VerificarTokens(tokens []string) bool {
	if len(tokens) < 2 {
		return false
	}
	for _, token := range tokens {
		if !VerificarToken(token) {
			return false
		}
	}
	return true
}
