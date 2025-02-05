package calculadora

import (
	"cd/verificaciones"
	"errors"
	"strconv"
	"tdas/pila"
)

type operacion struct {
	simbolo string
	aridad  int
	operar  func(numeros []int64) (int64, error)
}

var operaciones = []operacion{
	{simbolo: verificaciones.OP_SUMA, aridad: 2, operar: CalcularSuma},
	{simbolo: verificaciones.OP_RESTA, aridad: 2, operar: CalcularResta},
	{simbolo: verificaciones.OP_MULTIPLICACION, aridad: 2, operar: CalcularMultiplicacion},
	{simbolo: verificaciones.OP_DIVISION, aridad: 2, operar: CalcularDivision},
	{simbolo: verificaciones.OP_RAIZ_CUADRADA, aridad: 1, operar: CalcularRaizCuadrada},
	{simbolo: verificaciones.OP_EXPONENCIACION, aridad: 2, operar: CalcularExponenciacion},
	{simbolo: verificaciones.OP_LOGARITMO, aridad: 2, operar: CalcularLogaritmo},
	{simbolo: verificaciones.OP_TERNARIO, aridad: 3, operar: CalcularTernario},
}

func Calculadora(tokens []string) (int64, error) {
	numeros := pila.CrearPilaDinamica[int64]()
	if !verificaciones.VerificarTokens(tokens) {
		return 0, errors.New("token invalido detectado")
	}
	for _, token := range tokens {
		numero, err := strconv.ParseInt(token, 10, 64)
		if err == nil {
			numeros.Apilar(numero)
			continue
		}

		for _, operacion := range operaciones {
			if operacion.simbolo == token {
				operandos := make([]int64, operacion.aridad)
				for i := operacion.aridad - 1; i >= 0; i-- {
					operando, err := desapilarConVerificacion(numeros)
					if err != nil {
						return 0, err
					}
					operandos[i] = operando
				}
				resultado, err := operacion.operar(operandos)
				if err != nil {
					return 0, err
				}
				numeros.Apilar(resultado)
				break
			}
		}

	}
	resultado := numeros.Desapilar()
	if numeros.EstaVacia() {
		return resultado, nil
	}
	return 0, errors.New("no se realizaron suficientes operaciones")
}
