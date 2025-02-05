package expulsion

import (
	"errors"
	"fmt"
	"strconv"
	"tdas/cola_prioridad"
	"tdas/diccionario"
	"tp2/auxiliares"
)

type recurso struct {
	nombre   string
	cantidad int
}

func ImprimirSospechososEnOrden(ipsSospechosas []string) {
	if len(ipsSospechosas) > 0 {
		colaAux := cola_prioridad.CrearHeapArr(ipsSospechosas, auxiliares.CompararIpsMin)
		for !colaAux.EstaVacia() {
			fmt.Printf("DoS: %s\n", colaAux.Desencolar())
		}
	}
	fmt.Println("OK")
}

// Caso promedio O(log v) Peor Caso O(v) siendo v la cantidad de visitantes
func VerVisitantes(visitantes diccionario.DiccionarioOrdenado[string, bool], desde string, hasta string) error {
	if !auxiliares.EsIp(desde) || !auxiliares.EsIp(hasta) {
		return errors.New("parametros invalidos")
	}
	if visitantes.Cantidad() > 0 {
		fmt.Println("Visitantes:")
		iterador := visitantes.IteradorRango(&desde, &hasta)
		for iterador.HaySiguiente() {
			ip, _ := iterador.VerActual()
			fmt.Printf("\t%s\n", ip)
			iterador.Siguiente()
		}
	}
	fmt.Println("OK")
	return nil
}

// O(s + k log s) siendo s la cantidad de recursos y k el parametro
func VerMasVisitados(recursos diccionario.Diccionario[string, int], kStr string) error {
	k, err := strconv.Atoi(kStr)
	if err != nil || k < 0 {
		return err
	}
	if k == 0 {
		return nil
	}

	recursosArr := []recurso{}
	iterador := recursos.Iterador()
	for iterador.HaySiguiente() {
		nombreRecurso, cantidad := iterador.VerActual()
		recursoYCantidad := recurso{nombre: nombreRecurso, cantidad: cantidad}
		recursosArr = append(recursosArr, recursoYCantidad)
		iterador.Siguiente()
	}

	recursosTop := cola_prioridad.CrearHeapArr(recursosArr, func(recurso1, recurso2 recurso) int {
		return recurso1.cantidad - recurso2.cantidad
	})

	if !recursosTop.EstaVacia() {
		fmt.Println("Sitios más visitados:")
		for i := 0; i < k && !recursosTop.EstaVacia(); i++ {
			recursoInfo := recursosTop.Desencolar()
			fmt.Printf("\t%s - %d\n", recursoInfo.nombre, recursoInfo.cantidad)
		}
	}
	fmt.Println("OK")
	return nil
}
