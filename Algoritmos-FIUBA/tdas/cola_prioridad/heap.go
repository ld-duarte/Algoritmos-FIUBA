package cola_prioridad

const tamInicial = 32
const factorRedimension = 2
const condicionRedimension = 4

type colaConPrioridad[T any] struct {
	datos []T
	cant  int
	cmp   func(T, T) int
}

func CrearHeap[T any](funcion_cmp func(T, T) int) ColaPrioridad[T] {
	datos := make([]T, tamInicial)
	return &colaConPrioridad[T]{datos: datos, cant: 0, cmp: funcion_cmp}
}

func CrearHeapArr[T any](arreglo []T, funcion_cmp func(T, T) int) ColaPrioridad[T] {
	datos := make([]T, len(arreglo)+tamInicial)
	copy(datos, arreglo)
	cola := &colaConPrioridad[T]{datos: datos, cant: len(arreglo), cmp: funcion_cmp}
	for i := cola.cant - 1; i >= 0; i-- {
		downHeap(cola.datos, i, cola.cant, cola.cmp)
	}
	return cola
}

func HeapSort[T any](elementos []T, funcion_cmp func(T, T) int) {
	for i := len(elementos) - 1; i >= 0; i-- {
		downHeap(elementos, i, len(elementos), funcion_cmp)
	}
	for i := len(elementos) - 1; i > 0; i-- {
		elementos[0], elementos[i] = elementos[i], elementos[0]
		downHeap(elementos, 0, i, funcion_cmp)
	}
}

func (cola *colaConPrioridad[T]) EstaVacia() bool {
	return cola.cant == 0
}

func upHeap[T any](arreglo []T, pos int, funcion_cmp func(T, T) int) {
	for pos > 0 {
		padre := (pos - 1) / 2
		if funcion_cmp(arreglo[pos], arreglo[padre]) <= 0 {
			return
		}
		arreglo[pos], arreglo[padre] = arreglo[padre], arreglo[pos]
		pos = padre
	}
}

func (cola *colaConPrioridad[T]) redimensionar(nuevaCapacidad int) {
	nuevosDatos := make([]T, nuevaCapacidad)
	copy(nuevosDatos, cola.datos[:cola.cant])
	cola.datos = nuevosDatos
}

func (cola *colaConPrioridad[T]) Encolar(dato T) {
	if cola.cant == cap(cola.datos) {
		nuevaCapacidad := cap(cola.datos) * factorRedimension
		cola.redimensionar(nuevaCapacidad)
	}
	cola.datos[cola.cant] = dato
	cola.cant++
	upHeap(cola.datos, cola.cant-1, cola.cmp)
}

func (cola *colaConPrioridad[T]) entrarEnPanico() {
	if cola.EstaVacia() {
		panic("La cola esta vacia")
	}
}

func (cola *colaConPrioridad[T]) VerMax() T {
	cola.entrarEnPanico()
	return cola.datos[0]
}

func downHeap[T any](arreglo []T, pos int, cantidad int, funcion_cmp func(T, T) int) {
	for pos < cantidad {
		hijoIzq := 2*pos + 1
		hijoDer := 2*pos + 2
		mayor := pos
		if hijoIzq < cantidad && funcion_cmp(arreglo[hijoIzq], arreglo[mayor]) > 0 {
			mayor = hijoIzq
		}
		if hijoDer < cantidad && funcion_cmp(arreglo[hijoDer], arreglo[mayor]) > 0 {
			mayor = hijoDer
		}
		if mayor == pos {
			return
		}
		arreglo[pos], arreglo[mayor] = arreglo[mayor], arreglo[pos]
		pos = mayor
	}
}

func (cola *colaConPrioridad[T]) Desencolar() T {
	cola.entrarEnPanico()
	dato := cola.datos[0]
	cola.cant--
	if cola.cant > 0 {
		cola.datos[0] = cola.datos[cola.cant]
		downHeap(cola.datos, 0, cola.cant, cola.cmp)
	}
	if cola.cant > 0 && cola.cant <= cap(cola.datos)/condicionRedimension {
		nuevaCapacidad := cap(cola.datos) / factorRedimension
		if nuevaCapacidad >= tamInicial {
			cola.redimensionar(nuevaCapacidad)
		}
	}
	return dato
}

func (cola *colaConPrioridad[T]) Cantidad() int {
	return cola.cant
}
