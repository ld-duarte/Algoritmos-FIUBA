#!/usr/bin/python3

import funciones, csv, sys, rdfy_constantes, rdfy_comandos, rdfy_grafos


def main():
    if len(sys.argv) < 2:
        print("Error: Se debe proporcionar un archivo TSV.")
        sys.exit(1)

    archivo = sys.argv[1]
    datos = leer_archivo(archivo)
    usuario_cancion, grupos = rdfy_grafos.crear_grafo_gustos(datos)
    grafo_canciones, pagerank = None, None

    for linea in sys.stdin:
        linea = linea.strip()
        linea_dividida = linea.split(" ", 1)

        if len(linea_dividida) != 2: 
            continue
        
        comando, parametros = linea_dividida
        
        if comando not in rdfy_constantes.COMANDOS:
            print("Comando invalido")
            continue
        
        if comando in {rdfy_constantes.COMANDO_CICLO, rdfy_constantes.COMANDO_RANGO} and not grafo_canciones:
            grafo_canciones = rdfy_grafos.crear_grafo_canciones(datos)
        
        if comando == rdfy_constantes.COMANDO_IMPORTANTES and not pagerank:
            pagerank = funciones.page_rank(usuario_cancion)
        
        rdfy_comandos.ejecutar_comando(comando, parametros, usuario_cancion, grafo_canciones, pagerank, grupos)


def leer_archivo(archivo_tsv):
    datos = []
    with open(archivo_tsv, "r", encoding = "utf-8") as archivo:
        lector = csv.DictReader(archivo, delimiter = "\t")
        for linea in lector:
            datos.append(linea)
    return datos


if __name__ == "__main__":
    main()

       