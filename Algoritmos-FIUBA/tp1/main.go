package main

import (
	"bufio"
	"cd/calculadora"
	"fmt"
	"os"
	"strings"
)

func procesarEntradas(entrada *os.File) {
	scanner := bufio.NewScanner(entrada)
	for scanner.Scan() {
		linea := scanner.Text()
		tokens := strings.Fields(linea)
		resultado, err := calculadora.Calculadora(tokens)
		if err == nil {
			fmt.Println(resultado)
		} else {
			fmt.Println("ERROR")
		}
	}
}

func main() {
	procesarEntradas(os.Stdin)
}
