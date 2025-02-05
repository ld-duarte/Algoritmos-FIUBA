package cola

type nodoCola[T any] struct {
	elemento  T
	siguiente *nodoCola[T]
}

type colaEnlazada[T any] struct {
	primero *nodoCola[T]
	ultimo  *nodoCola[T]
}

func CrearColaEnlazada[T any]() Cola[T] {
	return &colaEnlazada[T]{}
}

func (cola *colaEnlazada[T]) EstaVacia() bool {
	return cola.primero == nil
}

func (cola *colaEnlazada[T]) verificarPanic() {
	if cola.EstaVacia() {
		panic("La cola esta vacia")
	}
}

func (cola *colaEnlazada[T]) VerPrimero() T {
	cola.verificarPanic()
	return cola.primero.elemento
}

func (cola *colaEnlazada[T]) crearNodoCola(dato T) *nodoCola[T] {
	return &nodoCola[T]{
		elemento:  dato,
		siguiente: nil,
	}
}
func (cola *colaEnlazada[T]) Encolar(dato T) {
	nuevoNodo := cola.crearNodoCola(dato)
	if cola.EstaVacia() {
		cola.primero = nuevoNodo
	} else {
		cola.ultimo.siguiente = nuevoNodo
	}
	cola.ultimo = nuevoNodo
}

func (cola *colaEnlazada[T]) Desencolar() T {
	cola.verificarPanic()
	valorDesencolado := cola.primero.elemento
	cola.primero = cola.primero.siguiente
	if cola.primero == nil {
		cola.ultimo = nil
	}
	return valorDesencolado
}
