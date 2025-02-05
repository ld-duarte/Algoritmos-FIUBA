package diccionario

import (
	TDAPila "tdas/pila"
)

type nodoAbb[K comparable, V any] struct {
	izquierdo *nodoAbb[K, V]
	derecho   *nodoAbb[K, V]
	clave     K
	dato      V
}

type abb[K comparable, V any] struct {
	raiz     *nodoAbb[K, V]
	cantidad int
	cmp      func(K, K) int
}

func CrearABB[K comparable, V any](funcion_cmp func(K, K) int) DiccionarioOrdenado[K, V] {
	abb := new(abb[K, V])
	abb.raiz = nil
	abb.cantidad = 0
	abb.cmp = funcion_cmp
	return abb
}

func (abb *abb[K, V]) crearNodo(clave K, dato V) *nodoAbb[K, V] {
	return &nodoAbb[K, V]{clave: clave, dato: dato}
}

func (abb *abb[K, V]) buscarClaveConPadre(nodo *nodoAbb[K, V], clave K) (*nodoAbb[K, V], *nodoAbb[K, V]) {
	var padre *nodoAbb[K, V] = nil
	for nodo != nil {
		comparacion := abb.cmp(clave, nodo.clave)
		if comparacion < 0 {
			padre = nodo
			nodo = nodo.izquierdo
		} else if comparacion > 0 {
			padre = nodo
			nodo = nodo.derecho
		} else {
			return nodo, padre
		}
	}
	return nil, padre
}

func (abb *abb[K, V]) guardarNodo(nodo *nodoAbb[K, V], clave K, dato V) *nodoAbb[K, V] {
	if nodo == nil {
		abb.cantidad++
		return abb.crearNodo(clave, dato)
	}

	comparacion := abb.cmp(clave, nodo.clave)
	if comparacion < 0 {
		nodo.izquierdo = abb.guardarNodo(nodo.izquierdo, clave, dato)
	} else if comparacion > 0 {
		nodo.derecho = abb.guardarNodo(nodo.derecho, clave, dato)
	} else {
		nodo.dato = dato
	}
	return nodo
}

func (abb *abb[K, V]) Guardar(clave K, dato V) {
	if abb.raiz != nil {
		nodo, _ := abb.buscarClaveConPadre(abb.raiz, clave)
		if nodo != nil {
			nodo.dato = dato
			return
		}
	}
	abb.raiz = abb.guardarNodo(abb.raiz, clave, dato)
}

func (abb *abb[K, V]) Pertenece(clave K) bool {
	nodo, _ := abb.buscarClaveConPadre(abb.raiz, clave)
	return nodo != nil
}

func (abb *abb[K, V]) verificarPanic(nodo *nodoAbb[K, V]) {
	if nodo == nil {
		panic("La clave no pertenece al diccionario")
	}
}

func (abb *abb[K, V]) Obtener(clave K) V {
	nodo, _ := abb.buscarClaveConPadre(abb.raiz, clave)
	abb.verificarPanic(nodo)
	return nodo.dato
}

func (abb *abb[K, V]) encontrarSucesor(nodo *nodoAbb[K, V]) *nodoAbb[K, V] {
	sucesor := nodo
	for sucesor.izquierdo != nil {
		sucesor = sucesor.izquierdo
	}
	return sucesor
}

func (abb *abb[K, V]) Borrar(clave K) V {
	nodo, _ := abb.buscarClaveConPadre(abb.raiz, clave)
	abb.verificarPanic(nodo)
	dato := nodo.dato
	abb.raiz = abb.borrarNodoConPadre(abb.raiz, nodo, nil)
	return dato
}

func (abb *abb[K, V]) borrarNodoConPadre(nodo, aBorrar, _ *nodoAbb[K, V]) *nodoAbb[K, V] {
	if nodo == nil {
		return nil
	}
	comparacion := abb.cmp(aBorrar.clave, nodo.clave)
	if comparacion < 0 {
		nodo.izquierdo = abb.borrarNodoConPadre(nodo.izquierdo, aBorrar, nodo)
	} else if comparacion > 0 {
		nodo.derecho = abb.borrarNodoConPadre(nodo.derecho, aBorrar, nodo)
	} else {
		if nodo.izquierdo != nil && nodo.derecho != nil {
			nodoSucesor := abb.encontrarSucesor(nodo.derecho)
			nodo.clave = nodoSucesor.clave
			nodo.dato = nodoSucesor.dato
			nodo.derecho = abb.borrarNodoConPadre(nodo.derecho, nodoSucesor, nodo)
		} else {
			abb.cantidad--
			if nodo.izquierdo != nil {
				return nodo.izquierdo
			}
			return nodo.derecho
		}
	}
	return nodo
}

func (abb *abb[K, V]) Cantidad() int {
	return abb.cantidad
}

func (abb *abb[K, V]) iterarRango(nodo *nodoAbb[K, V], desde *K, hasta *K, visitar func(clave K, dato V) bool) bool {
	if nodo == nil {
		return true
	}
	if desde == nil || abb.cmp(nodo.clave, *desde) >= 0 {
		if !abb.iterarRango(nodo.izquierdo, desde, hasta, visitar) {
			return false
		}
	}
	if (desde == nil || abb.cmp(nodo.clave, *desde) >= 0) && (hasta == nil || abb.cmp(nodo.clave, *hasta) <= 0) {
		if !visitar(nodo.clave, nodo.dato) {
			return false
		}
	}
	if hasta == nil || abb.cmp(nodo.clave, *hasta) <= 0 {
		if !abb.iterarRango(nodo.derecho, desde, hasta, visitar) {
			return false
		}
	}
	return true
}

func (abb *abb[K, V]) IterarRango(desde *K, hasta *K, visitar func(clave K, dato V) bool) {
	abb.iterarRango(abb.raiz, desde, hasta, visitar)
}

func (abb *abb[K, V]) Iterar(visitar func(clave K, dato V) bool) {
	abb.iterarRango(abb.raiz, nil, nil, visitar)
}

type iterAbb[K comparable, V any] struct {
	abb   *abb[K, V]
	pila  TDAPila.Pila[*nodoAbb[K, V]]
	desde *K
	hasta *K
}

func (iterador *iterAbb[K, V]) apilarEnOrden(nodo *nodoAbb[K, V], desde *K, hasta *K) {
	for nodo != nil {
		if (desde == nil || iterador.abb.cmp(nodo.clave, *desde) >= 0) && (hasta == nil || iterador.abb.cmp(nodo.clave, *hasta) <= 0) {
			iterador.pila.Apilar(nodo)
		}
		if desde != nil && iterador.abb.cmp(nodo.clave, *desde) < 0 {
			nodo = nodo.derecho
		} else {
			nodo = nodo.izquierdo
		}
	}
}

func (abb *abb[K, V]) IteradorRango(desde *K, hasta *K) IterDiccionario[K, V] {
	iterador := new(iterAbb[K, V])
	iterador.abb = abb
	iterador.pila = TDAPila.CrearPilaDinamica[*nodoAbb[K, V]]()
	iterador.apilarEnOrden(abb.raiz, desde, hasta)
	iterador.desde, iterador.hasta = desde, hasta
	return iterador
}

func (abb *abb[K, V]) Iterador() IterDiccionario[K, V] {
	return abb.IteradorRango(nil, nil)
}

func (iterador *iterAbb[K, V]) HaySiguiente() bool {
	return !iterador.pila.EstaVacia()
}

func (iterador *iterAbb[K, V]) VerActual() (K, V) {
	if !iterador.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	actual := iterador.pila.VerTope()
	return actual.clave, actual.dato
}

func (iterador *iterAbb[K, V]) Siguiente() {
	if !iterador.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	desapilado := iterador.pila.Desapilar()
	if desapilado != nil {
		hijoDer := desapilado.derecho
		if hijoDer != nil {
			iterador.apilarEnOrden(hijoDer, iterador.desde, iterador.hasta)
		}
	}
}
