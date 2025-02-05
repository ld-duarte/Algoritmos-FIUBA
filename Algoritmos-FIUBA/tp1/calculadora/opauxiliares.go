package calculadora

import (
	"errors"
	"math"
	"tdas/pila"
)

func desapilarConVerificacion(numeros pila.Pila[int64]) (int64, error) {
	if numeros.EstaVacia() {
		return 0, errors.New("cantidad de operandos insuficiente")
	}
	return numeros.Desapilar(), nil
}

func CalcularSuma(numeros []int64) (int64, error) {
	a, b := numeros[0], numeros[1]
	return a + b, nil
}

func CalcularResta(numeros []int64) (int64, error) {
	a, b := numeros[0], numeros[1]
	return a - b, nil
}

func CalcularMultiplicacion(numeros []int64) (int64, error) {
	a, b := numeros[0], numeros[1]
	return a * b, nil
}

func CalcularDivision(numeros []int64) (int64, error) {
	a, b := numeros[0], numeros[1]
	if b == 0 {
		return 0, errors.New("division por 0")
	}
	return a / b, nil
}

func CalcularExponenciacion(numeros []int64) (int64, error) {
	a, b := numeros[0], numeros[1]
	if b < 0 {
		return 0, errors.New("potencia con exponente negativo")
	}
	return int64(math.Pow(float64(a), float64(b))), nil
}

func CalcularLogaritmo(numeros []int64) (int64, error) {
	a, b := numeros[0], numeros[1]
	if a <= 0 {
		return 0, errors.New("argumento del logaritmo menor o igual a 0")
	} else if b <= 1 {
		return 0, errors.New("base del logaritmo menor a 2")
	}
	return int64(math.Log(float64(a)) / math.Log(float64(b))), nil
}

func CalcularRaizCuadrada(numeros []int64) (int64, error) {
	a := numeros[0]
	if a < 0 {
		return 0, errors.New("raiz de un numero negativo")
	}
	resultado := int64(math.Sqrt(float64(a)))
	return resultado, nil
}

func CalcularTernario(numeros []int64) (int64, error) {
	a, b, c := numeros[0], numeros[1], numeros[2]
	if a != 0 {
		return b, nil
	}
	return c, nil
}
