package pila_test

import (
	TDAPila "tdas/pila"
	"testing"

	"github.com/stretchr/testify/require"
)

const _ITERACIONES_VOLUMEN = 10000
const _ITERACIONES = 50

func TestPilaVacia(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	require.True(t, pila.EstaVacia())
	require.Panics(t, func() { pila.Desapilar() })
	require.Panics(t, func() { pila.VerTope() })
}

func TestPilaVerTope(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	for i := 0; i < _ITERACIONES; i++ {
		pila.Apilar(i)
		require.Equal(t, i, pila.VerTope())
	}
	for i := _ITERACIONES - 1; i >= 0; i-- {
		require.Equal(t, i, pila.VerTope())
		pila.Desapilar()
	}
}

func TestPilaApilar(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[string]()
	pila.Apilar("1")
	require.Equal(t, "1", pila.VerTope())
	pila.Apilar("27")
	require.Equal(t, "27", pila.VerTope())
}

func TestPilaDesapilar(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	for i := 0; i < _ITERACIONES; i++ {
		pila.Apilar(i)
		require.Equal(t, i, pila.VerTope())
	}
	for i := _ITERACIONES - 1; i >= 0; i-- {
		require.Equal(t, i, pila.Desapilar())

	}
}

func TestPilaVolumen(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	for i := 0; i < _ITERACIONES_VOLUMEN; i++ {
		pila.Apilar(i)
		require.Equal(t, i, pila.VerTope())
	}
	for i := _ITERACIONES_VOLUMEN - 1; i >= 0; i-- {
		require.Equal(t, i, pila.VerTope())
		require.Equal(t, i, pila.Desapilar())
	}
}

func TestPilaVaciada(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	pila.Apilar(1)
	require.False(t, pila.EstaVacia())
	pila.Apilar(2)
	pila.Desapilar()
	pila.Desapilar()
	require.True(t, pila.EstaVacia())
	require.Panics(t, func() { pila.Desapilar() })
	require.Panics(t, func() { pila.VerTope() })
}

func TestPilaPrimitivos(t *testing.T) {
	pilaString := TDAPila.CrearPilaDinamica[string]()
	pilaBool := TDAPila.CrearPilaDinamica[bool]()
	pilaRuna := TDAPila.CrearPilaDinamica[rune]()

	pilaString.Apilar("a")
	pilaString.Apilar("b")
	require.Equal(t, "b", pilaString.Desapilar())
	require.Equal(t, "a", pilaString.VerTope())
	require.False(t, pilaString.EstaVacia())

	pilaBool.Apilar(true)
	pilaBool.Apilar(false)
	require.Equal(t, false, pilaBool.Desapilar())
	require.Equal(t, true, pilaBool.VerTope())
	require.False(t, pilaBool.EstaVacia())

	pilaRuna.Apilar('😄')
	require.Equal(t, '😄', pilaRuna.VerTope())
	require.Equal(t, '😄', pilaRuna.Desapilar())
	require.True(t, pilaRuna.EstaVacia())

}

type colores struct {
	color   string
	meGusta bool
}

func TestPilaStructs(t *testing.T) {
	pilaColores := TDAPila.CrearPilaDinamica[colores]()

	color1 := colores{"negro", true}
	color2 := colores{"verde", false}

	pilaColores.Apilar(color1)
	pilaColores.Apilar(color2)

	require.Equal(t, color2, pilaColores.Desapilar())
	require.Equal(t, color1, pilaColores.VerTope())
	require.False(t, pilaColores.EstaVacia())
}
