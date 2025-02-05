package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"tp0/ejercicios"
)

const _RUTA1 = "archivo1.in"
const _RUTA2 = "archivo2.in"

func guardarDatos(ruta string) ([]int, error) {
	arreglo := []int{}
	archivo, err := os.Open(ruta)
	if err != nil {
		return nil, err
	}
	defer archivo.Close()
	scanner1 := bufio.NewScanner(archivo)
	for scanner1.Scan() {
		numero, err := strconv.Atoi(scanner1.Text())
		if err != nil {
			return nil, err
		}
		arreglo = append(arreglo, numero)
	}
	return arreglo, err
}

func imprimirValores(arreglo []int) {
	for _, valor := range arreglo {
		fmt.Println(valor)
	}
}

func imprimirMayorOrdenado(arreglo1 []int, arreglo2 []int) {
	comparacion := ejercicios.Comparar(arreglo1, arreglo2)
	if comparacion == 1 {
		ejercicios.Seleccion(arreglo1)
		imprimirValores(arreglo1)
	} else if comparacion == -1 {
		ejercicios.Seleccion(arreglo2)
		imprimirValores(arreglo2)
	} else {
		fmt.Println("son iguales")
	}

}

func main() {
	arreglo1, err := guardarDatos(_RUTA1)
	if err != nil {
		fmt.Printf("Error al procesar %s: %v\n", _RUTA1, err)
		return
	}

	arreglo2, err := guardarDatos(_RUTA2)
	if err != nil {
		fmt.Printf("Error al procesar %s: %v\n", _RUTA2, err)
		return
	}

	imprimirMayorOrdenado(arreglo1, arreglo2)

}
