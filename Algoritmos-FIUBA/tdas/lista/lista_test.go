package lista_test

import (
	"tdas/lista"
	"testing"

	"github.com/stretchr/testify/require"
)

const ITERACIONES_VOLUMEN = 50000

// test para probar que se pueda crear una lista vacía y que se comporte como tal
func TestListaVacia(t *testing.T) {
	l := lista.CrearListaEnlazada[int]()
	if !l.EstaVacia() {
		t.Error("La lista debería estar vacía")
	}
}

// test para probar que se puedan insertar elementos en la lista y que se mantenga el invariante de la lista
func TestInsertar(t *testing.T) {
	l := lista.CrearListaEnlazada[int]()
	l.InsertarPrimero(1)
	l.InsertarPrimero(2)
	l.InsertarPrimero(3)
	if l.VerPrimero() != 3 {
		t.Error("El primer elemento debería ser 3")
	}
	if l.VerUltimo() != 1 {
		t.Error("El último elemento debería ser 1")
	}
}

// test para probar que se pueda borrar el primer elemento de la lista y que se mantenga el invariante de la lista
func TestBorrar(t *testing.T) {
	l := lista.CrearListaEnlazada[int]()
	l.InsertarPrimero(1)
	l.InsertarPrimero(2)
	l.InsertarPrimero(3)
	if l.BorrarPrimero() != 3 {
		t.Error("El primer elemento borrado debería ser 3")
	}
	if l.VerPrimero() != 2 {
		t.Error("El primer elemento ahora debería ser 2")
	}
}

// test para probar que se pueda ver el primer elemento de la lista
func TestVerPrimero(t *testing.T) {
	l := lista.CrearListaEnlazada[int]()
	l.InsertarPrimero(1)
	if l.VerPrimero() != 1 {
		t.Error("El primer elemento debería ser 1")
	}
}

// test para probar que se pueda ver el último elemento de la lista
func TestVerUltimo(t *testing.T) {
	l := lista.CrearListaEnlazada[int]()
	l.InsertarPrimero(1)
	if l.VerUltimo() != 1 {
		t.Error("El último elemento debería ser 1")
	}
}

// test para probar que se pueda ver el largo de la lista
func TestLargo(t *testing.T) {
	l := lista.CrearListaEnlazada[int]()
	l.InsertarPrimero(1)
	l.InsertarPrimero(2)
	l.InsertarPrimero(3)
	if l.Largo() != 3 {
		t.Error("El largo debería ser 3")
	}
}

// test para probar que se pueda iterar sobre la lista
func TestIterar(t *testing.T) {
	l := lista.CrearListaEnlazada[int]()
	l.InsertarPrimero(1)
	l.InsertarPrimero(2)
	l.InsertarPrimero(3)
	var suma int
	l.Iterar(func(e int) bool {
		suma += e
		return true
	})
	if suma != 6 {
		t.Error("La suma debería ser 6")
	}
}

// test para probar que se pueda crear un iterador para la lista
func TestIterador(t *testing.T) {
	l := lista.CrearListaEnlazada[int]()
	l.InsertarPrimero(1)
	l.InsertarPrimero(2)
	l.InsertarPrimero(3)
	iter := l.Iterador()
	if iter.VerActual() != 3 {
		t.Error("El primer elemento debería ser 3")
	}
}

// test para probar que se pueda ver el elemento actual del iterador
func TestVerActual(t *testing.T) {
	l := lista.CrearListaEnlazada[int]()
	l.InsertarPrimero(1)
	iter := l.Iterador()
	if iter.VerActual() != 1 {
		t.Error("El elemento actual debería ser 1")
	}
}

// test para probar que se pueda avanzar al siguiente elemento del iterador
func TestSiguiente(t *testing.T) {
	l := lista.CrearListaEnlazada[int]()
	l.InsertarPrimero(1)
	l.InsertarPrimero(2)
	iter := l.Iterador()
	iter.Siguiente()
	if iter.VerActual() != 1 {
		t.Error("El elemento actual debería ser 1")
	}
}

// test para probar que se pueda insertar un elemento en la posición actual del iterador
func TestInsertarConIterador(t *testing.T) {
	l := lista.CrearListaEnlazada[int]()
	l.InsertarPrimero(1)
	l.InsertarPrimero(2)
	iter := l.Iterador()
	iter.Insertar(3)
	if l.VerPrimero() != 3 {
		t.Error("El primer elemento debería ser 3")
	}
}

// test para probar que se pueda borrar el elemento actual del iterador
func TestBorrarConIterador(t *testing.T) {
	l := lista.CrearListaEnlazada[int]()
	l.InsertarPrimero(1)
	l.InsertarPrimero(2)
	iter := l.Iterador()
	if iter.Borrar() != 2 {
		t.Error("El elemento borrado debería ser 2")
	}
	if l.VerPrimero() != 1 {
		t.Error("El primer elemento debería ser 1")
	}
}

// test para probar que se pueda iterar sobre la lista con corte
func TestIterarConCorte(t *testing.T) {
	l := lista.CrearListaEnlazada[int]()
	l.InsertarPrimero(1)
	l.InsertarPrimero(2)
	l.InsertarPrimero(3)
	l.InsertarPrimero(4)
	suma := 0
	l.Iterar(func(e int) bool {
		if e == 2 {
			return false
		}
		suma += e
		return true
	})
	if suma != 7 {
		t.Error("La suma debería ser 7")
	}
}

// test para probar que se puedan insertar múltiples tipos de datos en la lista
func TestDiferentesTipos(t *testing.T) {
	l := lista.CrearListaEnlazada[interface{}]()
	l.InsertarPrimero(1)
	l.InsertarPrimero("hola")
	if l.VerPrimero() != "hola" {
		t.Error("El primer elemento debería ser 'hola'")
	}
}

// test para probar que se pueda insertar un volumen grande de elementos
func TestVolumen(t *testing.T) {
	l := lista.CrearListaEnlazada[int]()
	for i := 0; i < ITERACIONES_VOLUMEN; i++ {
		l.InsertarPrimero(i)
		require.Equal(t, l.VerPrimero(), i)
		require.Equal(t, l.Largo(), i+1)
	}
	for i := ITERACIONES_VOLUMEN - 1; i >= 0; i-- {
		if l.BorrarPrimero() != i {
			t.Error("El elemento debería ser", i)
		}
	}
	require.True(t, l.EstaVacia())
}

// Test para insertar un elemento al principio con el iterador
func TestInsertarAlPrincipioIterador(t *testing.T) {
	l := lista.CrearListaEnlazada[int]()
	l.InsertarPrimero(1)
	iter := l.Iterador()
	iter.Insertar(2)
	if l.VerPrimero() != 2 {
		t.Error("El primer elemento debería ser 2")
	}
	if l.Largo() != 2 {
		t.Error("El largo debería ser 2")
	}
}

// Test para insertar un elemento al final usando el iterador
func TestInsertarAlFinalIterador(t *testing.T) {
	l := lista.CrearListaEnlazada[int]()
	l.InsertarPrimero(1)
	iter := l.Iterador()
	iter.Siguiente() // Nos movemos al final
	iter.Insertar(2)
	if l.VerUltimo() != 2 {
		t.Error("El último elemento debería ser 2")
	}
}

// Test para insertar un elemento en el medio con el iterador
func TestInsertarEnMedioConIterador(t *testing.T) {
	l := lista.CrearListaEnlazada[int]()
	l.InsertarPrimero(1)
	l.InsertarPrimero(3)
	iter := l.Iterador()
	iter.Siguiente()
	iter.Insertar(2)
	if l.VerPrimero() != 3 {
		t.Error("El primer elemento debería ser 3")
	}
	if l.VerUltimo() != 1 {
		t.Error("El último elemento debería ser 1")
	}
}

// Test para remover el primer elemento usando el iterador
func TestRemoverPrimeroConIterador(t *testing.T) {
	l := lista.CrearListaEnlazada[int]()
	l.InsertarPrimero(1)
	l.InsertarPrimero(2)
	iter := l.Iterador()
	iter.Borrar()
	if l.VerPrimero() != 1 {
		t.Error("El primer elemento debería ser 1")
	}
}

// Test para remover el último elemento usando el iterador
func TestRemoverUltimoConIterador(t *testing.T) {
	l := lista.CrearListaEnlazada[int]()
	l.InsertarPrimero(1)
	l.InsertarPrimero(2)
	iter := l.Iterador()
	iter.Siguiente() // Nos movemos al último elemento
	iter.Borrar()    // Remueve el último elemento
	if l.VerUltimo() != 2 {
		t.Error("El último elemento debería ser 2")
	}
}

// Test para remover un elemento en el medio con el iterador
func TestRemoverEnMedioConIterador(t *testing.T) {
	l := lista.CrearListaEnlazada[int]()
	l.InsertarPrimero(1)
	l.InsertarPrimero(2)
	l.InsertarPrimero(3)
	iter := l.Iterador()
	iter.Siguiente()
	iter.Borrar()
	if l.VerPrimero() != 3 {
		t.Error("El primer elemento debería ser 3")
	}
	if l.VerUltimo() != 1 {
		t.Error("El último elemento debería ser 1")
	}
}

func TestVolumenConIterador(t *testing.T) {
	l := lista.CrearListaEnlazada[int]()
	iter := l.Iterador()
	for i := 0; i < ITERACIONES_VOLUMEN; i++ {
		iter.Insertar(i)
		require.Equal(t, i, iter.VerActual())
	}
	require.Equal(t, l.Largo(), ITERACIONES_VOLUMEN)
	require.False(t, l.EstaVacia())
	for i := 0; i < ITERACIONES_VOLUMEN; i++ {
		require.True(t, iter.HaySiguiente())
		iter.Siguiente()
	}
	require.Panics(t, iter.Siguiente)
	iter2 := l.Iterador()
	for i := ITERACIONES_VOLUMEN - 1; i >= 0; i-- {
		require.Equal(t, iter2.Borrar(), i)
	}
	require.False(t, iter2.HaySiguiente())
	require.True(t, l.EstaVacia())
}
