package diccionario

import (
	"fmt"
	"hash/fnv"
)

const (
	estadoVacio   = 0
	estadoOcupado = 1
	estadoBorrado = 2

	tamañoInicial       = 31
	maximoFactorDeCarga = 0.7
	factorDeRedimension = 2
)

type celdaHash[K comparable, V any] struct {
	clave  K
	dato   V
	estado int
}

type hashCerrado[K comparable, V any] struct {
	tabla    []celdaHash[K, V]
	cantidad int
	tam      int
	borrados int
}

func convertirABytes[K comparable](clave K) []byte {
	return []byte(fmt.Sprintf("%v", clave))
}

func fnv1aHash(clave []byte) uint64 {
	h := fnv.New64a()
	h.Write(clave)
	return h.Sum64()
}

func (hash *hashCerrado[K, V]) calcularPosicion(clave K, tam int) int {
	return int(fnv1aHash(convertirABytes(clave)) % uint64(tam))
}

func (hash *hashCerrado[K, V]) buscar(clave K) (int, bool) {
	pos := hash.calcularPosicion(clave, hash.tam)
	for {
		if hash.tabla[pos].estado == estadoOcupado && hash.tabla[pos].clave == clave {
			return pos, true
		}
		if hash.tabla[pos].estado == estadoVacio {
			return -1, false
		}
		pos = (pos + 1) % hash.tam
	}
}

func (hash *hashCerrado[K, V]) factorDeCarga() float64 {
	return (float64(hash.cantidad + hash.borrados)) / float64(hash.tam)
}

func esPrimo(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func siguientePrimo(n int) int {
	n++
	for !esPrimo(n) {
		n++
	}
	return n
}

func (hash *hashCerrado[K, V]) crearTabla(tam int) []celdaHash[K, V] {
	return make([]celdaHash[K, V], tam)
}

func (hash *hashCerrado[K, V]) insertarCelda(celda celdaHash[K, V], tabla []celdaHash[K, V], tam int) {
	pos := hash.calcularPosicion(celda.clave, tam)
	for tabla[pos].estado == estadoOcupado {
		pos = (pos + 1) % tam
	}
	tabla[pos] = celda
}

func (hash *hashCerrado[K, V]) redimensionar() {
	nuevoTam := siguientePrimo(hash.tam * factorDeRedimension)
	nuevaTabla := hash.crearTabla(nuevoTam)

	for _, celda := range hash.tabla {
		if celda.estado == estadoOcupado {
			hash.insertarCelda(celda, nuevaTabla, nuevoTam)
		}
	}
	hash.tabla = nuevaTabla
	hash.tam = nuevoTam
	hash.borrados = 0
}

func CrearHash[K comparable, V any]() Diccionario[K, V] {
	hash := &hashCerrado[K, V]{}
	hash.tabla = hash.crearTabla(tamañoInicial)
	hash.tam = tamañoInicial
	return hash
}

func (hash *hashCerrado[K, V]) Guardar(clave K, dato V) {
	if hash.factorDeCarga() >= maximoFactorDeCarga {
		hash.redimensionar()
	}
	pos, encontrado := hash.buscar(clave)
	if encontrado {
		hash.tabla[pos].dato = dato
		return
	}
	celda := celdaHash[K, V]{clave: clave, dato: dato, estado: estadoOcupado}
	hash.insertarCelda(celda, hash.tabla, hash.tam)
	hash.cantidad++
}

func (hash *hashCerrado[K, V]) Pertenece(clave K) bool {
	_, encontrado := hash.buscar(clave)
	return encontrado
}

func (hash *hashCerrado[K, V]) Obtener(clave K) V {
	pos, encontrado := hash.buscar(clave)
	if encontrado {
		return hash.tabla[pos].dato
	}
	panic("La clave no pertenece al diccionario")
}

func (hash *hashCerrado[K, V]) Borrar(clave K) V {
	pos, encontrado := hash.buscar(clave)
	if !encontrado {
		panic("La clave no pertenece al diccionario")
	}
	valor := hash.tabla[pos].dato
	hash.tabla[pos].estado = estadoBorrado
	hash.cantidad--
	hash.borrados++
	if hash.factorDeCarga() >= maximoFactorDeCarga {
		hash.redimensionar()
	}
	return valor
}

func (hash *hashCerrado[K, V]) Cantidad() int {
	return hash.cantidad
}

func (hash *hashCerrado[K, V]) Iterar(fn func(clave K, dato V) bool) {
	for _, celda := range hash.tabla {
		if celda.estado == estadoOcupado {
			if !fn(celda.clave, celda.dato) {
				return
			}
		}
	}
}

type iterHashCerrado[K comparable, V any] struct {
	hash      *hashCerrado[K, V]
	posActual int
}

func (hash *hashCerrado[K, V]) Iterador() IterDiccionario[K, V] {
	for i := 0; i < hash.tam; i++ {
		if hash.tabla[i].estado == estadoOcupado {
			return &iterHashCerrado[K, V]{hash: hash, posActual: i}
		}
	}
	return &iterHashCerrado[K, V]{hash: hash, posActual: hash.tam}
}

func (iter *iterHashCerrado[K, V]) HaySiguiente() bool {
	return iter.posActual != iter.hash.tam
}

func (iter *iterHashCerrado[K, V]) VerActual() (K, V) {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	return iter.hash.tabla[iter.posActual].clave, iter.hash.tabla[iter.posActual].dato
}

func (iter *iterHashCerrado[K, V]) Siguiente() {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	for i := iter.posActual + 1; i < iter.hash.tam; i++ {
		if iter.hash.tabla[i].estado == estadoOcupado {
			iter.posActual = i
			return
		}
	}
	iter.posActual = iter.hash.tam
}
