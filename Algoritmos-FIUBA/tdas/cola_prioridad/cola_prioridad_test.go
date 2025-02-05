package cola_prioridad_test

import (
	"tdas/cola_prioridad"
	"testing"
)

// Comparación para un Heap de Máximo
func CompararMaximo(a, b int) int {
	if a > b {
		return 1 // a es mayor que b
	} else if a < b {
		return -1 // a es menor que b
	}
	return 0 // a es igual a b
}

// Comparación para un Heap de Mínimo
func CompararMinimo(a, b int) int {
	if a < b {
		return 1 // a es menor que b
	} else if a > b {
		return -1 // a es mayor que b
	}
	return 0 // a es igual a b
}

// Prueba para Heap de Máximo
func TestHeapMaximo(t *testing.T) {
	heap := cola_prioridad.CrearHeap(CompararMaximo)

	// Encolar elementos
	heap.Encolar(10)
	heap.Encolar(20)
	heap.Encolar(5)

	// Desencolar y verificar el orden
	if dato := heap.Desencolar(); dato != 20 {
		t.Errorf("Esperado 20, obtenido %d", dato)
	}
	if dato := heap.Desencolar(); dato != 10 {
		t.Errorf("Esperado 10, obtenido %d", dato)
	}
	if dato := heap.Desencolar(); dato != 5 {
		t.Errorf("Esperado 5, obtenido %d", dato)
	}
}

// Prueba para Heap de Mínimo
func TestHeapMinimo(t *testing.T) {
	heap := cola_prioridad.CrearHeap(CompararMinimo)

	// Encolar elementos
	heap.Encolar(10)
	heap.Encolar(20)
	heap.Encolar(5)

	// Desencolar y verificar el orden
	if dato := heap.Desencolar(); dato != 5 {
		t.Errorf("Esperado 5, obtenido %d", dato)
	}
	if dato := heap.Desencolar(); dato != 10 {
		t.Errorf("Esperado 10, obtenido %d", dato)
	}
	if dato := heap.Desencolar(); dato != 20 {
		t.Errorf("Esperado 20, obtenido %d", dato)
	}
}

// Benchmark para Heap
func BenchmarkHeap(t *testing.B) {
	for i := 0; i < t.N; i++ {
		heap := cola_prioridad.CrearHeap(CompararMaximo)
		for j := 0; j < 12500; j++ {
			heap.Encolar(j)
		}
		for j := 0; j < 12500; j++ {
			heap.Desencolar()
		}
	}
}
