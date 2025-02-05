package lista

type nodoLista[T any] struct {
	dato      T
	siguiente *nodoLista[T]
}

type listaEnlazada[T any] struct {
	primero *nodoLista[T]
	ultimo  *nodoLista[T]
	largo   int
}

type iterListaEnlazada[T any] struct {
	actual   *nodoLista[T]
	anterior *nodoLista[T]
	lista    *listaEnlazada[T]
}

func CrearListaEnlazada[T any]() Lista[T] {
	return &listaEnlazada[T]{}
}

func (lista *listaEnlazada[T]) verificarPanicLista() {
	if lista.EstaVacia() {
		panic("La lista esta vacia")
	}
}

func (lista *listaEnlazada[T]) crearNodoLista(elemento T) *nodoLista[T] {
	return &nodoLista[T]{
		dato:      elemento,
		siguiente: nil,
	}
}

func (lista *listaEnlazada[T]) EstaVacia() bool {
	return lista.largo == 0
}

func (lista *listaEnlazada[T]) InsertarPrimero(elemento T) {
	nuevoNodo := lista.crearNodoLista(elemento)
	if lista.EstaVacia() {
		lista.ultimo = nuevoNodo
	}
	nuevoNodo.siguiente = lista.primero
	lista.primero = nuevoNodo
	lista.largo++

}

func (lista *listaEnlazada[T]) InsertarUltimo(elemento T) {
	nuevoNodo := lista.crearNodoLista(elemento)
	if lista.EstaVacia() {
		lista.primero = nuevoNodo
	} else {
		lista.ultimo.siguiente = nuevoNodo
	}
	lista.ultimo = nuevoNodo
	lista.largo++
}

func (lista *listaEnlazada[T]) BorrarPrimero() T {
	lista.verificarPanicLista()
	dato := lista.primero.dato
	lista.primero = lista.primero.siguiente
	if lista.primero == nil {
		lista.ultimo = nil
	}
	lista.largo--
	return dato
}

func (lista *listaEnlazada[T]) VerPrimero() T {
	lista.verificarPanicLista()
	return lista.primero.dato
}

func (lista *listaEnlazada[T]) VerUltimo() T {
	lista.verificarPanicLista()
	return lista.ultimo.dato
}

func (lista *listaEnlazada[T]) Largo() int {
	return lista.largo
}

func (lista *listaEnlazada[T]) Iterar(visitar func(T) bool) {
	for actual := lista.primero; actual != nil; actual = actual.siguiente {
		if !visitar(actual.dato) {
			return
		}
	}
}

func (lista *listaEnlazada[T]) Iterador() IteradorLista[T] {
	return &iterListaEnlazada[T]{
		actual:   lista.primero,
		anterior: nil,
		lista:    lista,
	}
}

func (iterador *iterListaEnlazada[T]) verificarPanicIterador() {
	if !iterador.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
}

func (iterador *iterListaEnlazada[T]) VerActual() T {
	iterador.verificarPanicIterador()
	return iterador.actual.dato
}

func (iterador *iterListaEnlazada[T]) HaySiguiente() bool {
	return iterador.actual != nil
}

func (iterador *iterListaEnlazada[T]) Siguiente() {
	iterador.verificarPanicIterador()
	iterador.anterior = iterador.actual
	iterador.actual = iterador.actual.siguiente
}

func (iterador *iterListaEnlazada[T]) Insertar(elemento T) {
	nuevoNodo := iterador.lista.crearNodoLista(elemento)
	nuevoNodo.siguiente = iterador.actual
	if iterador.anterior == nil {
		iterador.lista.primero = nuevoNodo
	} else {
		iterador.anterior.siguiente = nuevoNodo
	}
	if iterador.actual == nil {
		iterador.lista.ultimo = nuevoNodo
	}
	iterador.actual = nuevoNodo
	iterador.lista.largo++
}

func (iterador *iterListaEnlazada[T]) Borrar() T {
	iterador.verificarPanicIterador()
	valor := iterador.actual.dato
	if iterador.anterior == nil {
		iterador.lista.primero = iterador.actual.siguiente
	} else {
		iterador.anterior.siguiente = iterador.actual.siguiente
	}
	if iterador.actual.siguiente == nil {
		iterador.lista.ultimo = iterador.anterior
	}
	iterador.actual = iterador.actual.siguiente
	iterador.lista.largo--
	return valor
}
