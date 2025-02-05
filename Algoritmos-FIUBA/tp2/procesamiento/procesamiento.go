package procesamiento

import (
	"bufio"
	"os"
	"strings"
	"tdas/diccionario"
	"time"
	"tp2/expulsion"
)

const layoutISO = "2006-01-02T15:04:05-07:00"

func AgregarArchivo(nombreArchivo string, visitantes diccionario.DiccionarioOrdenado[string, bool], recursos diccionario.Diccionario[string, int]) error {
	file, err := os.Open(nombreArchivo)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	tiemposIps := diccionario.CrearHash[string, []time.Time]()
	ipsSospechosasYaVisitadas := diccionario.CrearHash[string, bool]() // para comprobar que ya se encuentre en el slice en O(1)
	ipsSospechosas := []string{}

	for scanner.Scan() {
		linea := scanner.Text()
		datos := strings.Split(linea, "\t")
		if len(datos) < 4 {
			continue
		}

		ip := datos[0]
		timestamp := datos[1]
		recurso := datos[3]

		tiempo, err := time.Parse(layoutISO, timestamp)
		if err != nil {
			continue
		}

		guardarRecursos(recursos, recurso)
		visitantes.Guardar(ip, true)
		guardarTiempos(tiemposIps, ip, tiempo)

		if esSospechosoDeDoS(tiemposIps.Obtener(ip)) && !ipsSospechosasYaVisitadas.Pertenece(ip) {
			ipsSospechosas = append(ipsSospechosas, ip)
			ipsSospechosasYaVisitadas.Guardar(ip, true)
		}
	}
	expulsion.ImprimirSospechososEnOrden(ipsSospechosas)
	return nil
}

func guardarRecursos(recursos diccionario.Diccionario[string, int], recurso string) {
	if !recursos.Pertenece(recurso) {
		recursos.Guardar(recurso, 0)
	}
	recursos.Guardar(recurso, recursos.Obtener(recurso)+1)
}

func guardarTiempos(tiemposIps diccionario.Diccionario[string, []time.Time], ip string, tiempo time.Time) {
	if !tiemposIps.Pertenece(ip) {
		tiemposIps.Guardar(ip, []time.Time{})
	}
	tiempos := tiemposIps.Obtener(ip)
	tiempos = append(tiempos, tiempo)
	tiemposIps.Guardar(ip, tiempos)
}

func esSospechosoDeDoS(tiempos []time.Time) bool {
	if len(tiempos) < 5 {
		return false
	}
	for i := 0; i <= len(tiempos)-5; i++ {
		if tiempos[i+4].Sub(tiempos[i]) < 2*time.Second {
			return true
		}
	}
	return false
}
