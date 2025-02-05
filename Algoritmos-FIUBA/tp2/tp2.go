package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
	"tdas/diccionario"
	"tp2/auxiliares"
	"tp2/expulsion"
	"tp2/procesamiento"
)

const (
	comandoAgAr      = "agregar_archivo"
	comandoVerVis    = "ver_visitantes"
	comandoVerMasVis = "ver_mas_visitados"
)

type comando struct {
	nombre     string
	parametros int
}

var comandos = []comando{
	{nombre: comandoAgAr, parametros: 2},
	{nombre: comandoVerVis, parametros: 3},
	{nombre: comandoVerMasVis, parametros: 2},
}

func main() {
	visitantes := diccionario.CrearABB[string, bool](auxiliares.CompararIpsMax)
	recursos := diccionario.CrearHash[string, int]()
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		linea := scanner.Text()
		parametros := strings.Fields(linea)
		if len(parametros) == 0 {
			continue
		}

		comandoEncontrado := false
		for _, comando := range comandos {
			if comando.nombre == parametros[0] && comando.parametros == len(parametros) {
				err := accionar(comando.nombre, parametros, visitantes, recursos)
				if err != nil {
					imprimirError(comando.nombre)
					return
				}
				comandoEncontrado = true
				break
			}
		}
		if !comandoEncontrado {
			imprimirError(parametros[0])
			return
		}
	}
}

func imprimirError(comando string) {
	fmt.Fprintf(os.Stderr, "Error en comando %s\n", comando)
}

func accionar(comando string, parametros []string, visitantes diccionario.DiccionarioOrdenado[string, bool], recursos diccionario.Diccionario[string, int]) error {
	switch comando {
	case comandoAgAr:
		return procesamiento.AgregarArchivo(parametros[1], visitantes, recursos)
	case comandoVerVis:
		return expulsion.VerVisitantes(visitantes, parametros[1], parametros[2])
	case comandoVerMasVis:
		return expulsion.VerMasVisitados(recursos, parametros[1])
	default:
		return errors.New("comando invalido")
	}
}
