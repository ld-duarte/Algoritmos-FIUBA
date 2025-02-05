package cola_test

import (
	TDACola "tdas/cola"
	"testing"

	"github.com/stretchr/testify/require"
)

const _ITERACIONES_VOLUMEN = 10000

func TestColaVacia(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()
	require.True(t, cola.EstaVacia())
	require.Panics(t, func() { cola.Desencolar() })
	require.Panics(t, func() { cola.VerPrimero() })
}

func TestColaVerPrimero(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[string]()
	cola.Encolar("buenas")
	cola.Encolar("milanesa")
	require.Equal(t, "buenas", cola.VerPrimero())
	cola.Desencolar()
	require.Equal(t, "milanesa", cola.VerPrimero())
}

func TestColaVolumen(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()
	for i := 0; i < _ITERACIONES_VOLUMEN; i++ {
		cola.Encolar(i)
	}
	for i := 0; i < _ITERACIONES_VOLUMEN; i++ {
		require.Equal(t, i, cola.Desencolar())
	}
	require.True(t, cola.EstaVacia())
	require.Panics(t, func() { cola.Desencolar() })
	require.Panics(t, func() { cola.VerPrimero() })
}

func TestColaPrimitivos(t *testing.T) {
	colaString := TDACola.CrearColaEnlazada[string]()
	colaBool := TDACola.CrearColaEnlazada[bool]()
	colaRuna := TDACola.CrearColaEnlazada[rune]()

	colaString.Encolar("a")
	colaString.Encolar("b")
	require.Equal(t, "a", colaString.Desencolar())
	require.Equal(t, "b", colaString.VerPrimero())
	require.False(t, colaString.EstaVacia())

	colaBool.Encolar(true)
	colaBool.Encolar(false)
	require.Equal(t, true, colaBool.Desencolar())
	require.Equal(t, false, colaBool.VerPrimero())
	require.False(t, colaBool.EstaVacia())

	colaRuna.Encolar('😄')
	require.Equal(t, '😄', colaRuna.VerPrimero())
	require.Equal(t, '😄', colaRuna.Desencolar())
	require.True(t, colaRuna.EstaVacia())
}

type casas struct {
	tipo    string
	vendida bool
}

func TestColaStructs(t *testing.T) {
	colaCasas := TDACola.CrearColaEnlazada[casas]()

	casa1 := casas{"moderna", true}
	casa2 := casas{"antigua", false}

	colaCasas.Encolar(casa1)
	colaCasas.Encolar(casa2)

	require.Equal(t, casa1, colaCasas.Desencolar())
	require.Equal(t, casa2, colaCasas.VerPrimero())
	require.False(t, colaCasas.EstaVacia())
}
