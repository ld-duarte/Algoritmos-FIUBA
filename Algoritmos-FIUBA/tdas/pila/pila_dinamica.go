package pila

const _TAM_INICIAL = 10
const _FACTOR_REDIMENSION = 2
const _CONDICION_DE_REDIMENSION = 4

/* Definición del struct pila proporcionado por la cátedra. */

type pilaDinamica[T any] struct {
	datos    []T
	cantidad int
}

func (pila *pilaDinamica[T]) verificarPanic() {
	if pila.EstaVacia() {
		panic("La pila esta vacia")
	}
}

func CrearPilaDinamica[T any]() Pila[T] {
	pila := new(pilaDinamica[T])
	pila.datos = make([]T, _TAM_INICIAL)
	return pila
}

func (pila *pilaDinamica[T]) EstaVacia() bool {
	return pila.cantidad == 0
}

func (pila *pilaDinamica[T]) VerTope() T {
	pila.verificarPanic()
	valor := pila.datos[pila.cantidad-1]
	return valor
}

func (pila *pilaDinamica[T]) redimensionar(nuevaCapacidad int) {
	nuevaData := make([]T, nuevaCapacidad)
	copy(nuevaData, pila.datos)
	pila.datos = nuevaData
}

func (pila *pilaDinamica[T]) Apilar(dato T) {
	if pila.cantidad == cap(pila.datos) {
		nuevaCapacidad := pila.cantidad * _FACTOR_REDIMENSION
		pila.redimensionar(nuevaCapacidad)
	}
	pila.datos[pila.cantidad] = dato
	pila.cantidad++
}

func (pila *pilaDinamica[T]) Desapilar() T {
	pila.verificarPanic()
	pila.cantidad--
	valor := pila.datos[pila.cantidad]

	if pila.cantidad > 0 && pila.cantidad == cap(pila.datos)/_CONDICION_DE_REDIMENSION {
		nuevaCapacidad := cap(pila.datos) / _FACTOR_REDIMENSION
		pila.redimensionar(nuevaCapacidad)
	}
	return valor
}
