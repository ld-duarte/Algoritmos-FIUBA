package lista

type Lista[T any] interface {

	// EstaVacia devuelve true si la lista no tiene elementos, false en caso contrario.
	EstaVacia() bool

	// InsertarPrimero agrega un nuevo elemento al principio de la lista.
	InsertarPrimero(T)

	// InsertarUltimo agrega un nuevo elemento al final de la lista.
	InsertarUltimo(T)

	// BorrarPrimero borra el primer elemento de la lista. Si la lista tiene elementos, se borra el primer
	// elemento y se devuelve su valor. Si esta vacía, entra en pánico con un mensaje "La lista esta vacia".
	BorrarPrimero() T

	// VerPrimero obtiene el valor del primer elemento de la lista. Si la lista tiene elementos se devuelve
	// el valor del primero. Si esta vacía, entra en pánico con un mensaje "La lista esta vacia".
	VerPrimero() T

	// VerUltimo obtiene el valor del último elemento de la lista. Si la lista tiene elementos se devuelve
	// el valor del último. Si esta vacía, entra en pánico con un mensaje "La lista esta vacia".
	VerUltimo() T

	// Largo devuelve la cantidad de elementos de la lista.
	Largo() int

	// Iterar recorre los elementos de la lista aplicando la función visitar a cada uno. La iteracion continua
	// hasta que la lista se termine o hasta que la funcion visitar devuelva false.
	Iterar(visitar func(T) bool)

	//Iterador devuelve un iterador para recorrer los elementos de la lista manualmente.
	Iterador() IteradorLista[T]
}

type IteradorLista[T any] interface {

	// VerActual Devuelve el valor del elemento en la posición actual del iterador. Si se invoca sobre un
	// iterador que ya haya iterado todos los elementos, entra en pánico con un mensaje "El iterador
	// termino de iterar".
	VerActual() T

	// HaySiguiente Devuelve true si hay un siguiente elemento en la lista, false en caso contrario.
	HaySiguiente() bool

	// Siguiente avanza el iterador al siguiente elemento de la lista. Si se invoca sobre un iterador que ya
	// haya iterado todos los elementos, entra en pánico con un mensaje "El iterador termino de iterar".
	Siguiente()

	// Insertar agrega un elemento en la posición actual del iterador.
	Insertar(T)

	// Borrar elimina el elemento en la posición actual del iterador y devuelve su valor. Si se invoca sobre
	// un iterador que ya haya iterado todos los elementos, entra en pánico con un mensaje "El iterador
	// termino de iterar".
	Borrar() T
}
