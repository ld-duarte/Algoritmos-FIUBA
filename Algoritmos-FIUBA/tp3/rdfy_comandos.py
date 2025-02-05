import funciones, rdfy_constantes, heapq


##### EJECUCION #####
def ejecutar_comando(comando, parametros, usuario_cancion, grafo_canciones, pagerank, grupos):
    if comando == rdfy_constantes.COMANDO_CAMINO:
        parametros = parametros.split(rdfy_constantes.SEPARADOR_CANCIONES)
        camino_mas_corto(usuario_cancion, parametros[0], parametros[1], grupos)
    elif comando == rdfy_constantes.COMANDO_IMPORTANTES:
        cantidad = int(parametros)
        canciones_importantes(pagerank, cantidad, grupos)
    elif comando == rdfy_constantes.COMANDO_RECOMENDACION:
        parametros = parametros.split(" ", 2)
        canciones = parametros[2].split(rdfy_constantes.SEPARADOR_CANCIONES)
        recomendacion(usuario_cancion, parametros[0], int(parametros[1]), canciones, grupos)
    elif comando == rdfy_constantes.COMANDO_CICLO:
        parametros = parametros.split(" ", 1)
        n, cancion = int(parametros[0]), parametros[1]
        ciclo_canciones(grafo_canciones, n, cancion, grupos)
    elif comando == rdfy_constantes.COMANDO_RANGO:
        parametros = parametros.split(" ", 1)
        n, cancion = int(parametros[0]), parametros[1]
        rango_canciones(grafo_canciones, n, cancion, grupos)


##### COMANDO CAMINO #####
def camino_mas_corto(grafo, origen, destino, grupos):
    if origen not in grupos or destino not in grupos or grupos[origen] != rdfy_constantes.GRUPO_CANCIONES or grupos[destino] != rdfy_constantes.GRUPO_CANCIONES:
        print("Tanto el origen como el destino deben ser canciones")
        return
    padres = funciones.camino_minimo_bfs(grafo, destino, origen)

    if destino not in padres:
        print("No se encontro recorrido")
        return
    camino = reconstruir_camino_con_playlists(grafo, padres, destino)
    imprimir_camino_con_formato(camino, grupos)

def reconstruir_camino_con_playlists(grafo, padres, destino):
    camino_invertido = []
    actual = destino
    while actual is not None:
        camino_invertido.append(actual)
        if padres[actual] != None:
            camino_invertido.append(grafo.obtener_peso_arista(actual, padres[actual]))
        actual = padres[actual]
    return camino_invertido[::-1]

def imprimir_camino_con_formato(camino, grupos):
    for i in range(len(camino)-1):
        actual = camino[i]
        siguiente = camino[i+1]
        if grupos[actual] == rdfy_constantes.GRUPO_CANCIONES:
            print(f"{actual} --> aparece en playlist --> {siguiente} --> de -->", end=" ")
        if grupos[actual] == rdfy_constantes.GRUPO_USUARIOS:
            print(f"{actual} --> tiene una playlist --> {siguiente} --> donde aparece -->", end=" ")
    print(camino[-1])


##### TOP-K #####
def top_k(diccionario, grupo, cantidad, grupos):
    heap = []
    for vertice, valor in diccionario.items():
        if grupos[vertice] != grupo:
            continue

        if len(heap) < cantidad:
            heapq.heappush(heap,(valor, vertice))
        else: 
            importancia_desencolada, _= heap[0]
            if valor > importancia_desencolada:
                heapq.heappop(heap)
                heapq.heappush(heap, (valor, vertice))
    return heap


##### COMANDO IMPORTANTES #####
def canciones_importantes(pagerank, cantidad, grupos):    
    heap = top_k(pagerank, rdfy_constantes.GRUPO_CANCIONES, cantidad, grupos)
    top_canciones = []
    while heap:
        _, vertice = heapq.heappop(heap)
        top_canciones.append(vertice)
    print("; ".join(top_canciones[::-1]))


##### COMANDO RECOMENDACION #####
def recomendacion(grafo, grupo, cantidad, gustos, grupos):    
    similaridad_pr = funciones.page_rank_personalizado(grafo, gustos)
    heap = top_k(similaridad_pr, grupo, cantidad, grupos)
    recomendaciones = []
    while heap:
        _, recomendacion = heapq.heappop(heap)
        recomendaciones.append(recomendacion)
    print("; ".join(recomendaciones[::-1]))


##### COMANDO CICLO #####
def ciclo_canciones(grafo, largo, cancion, grupos):
    if grafo.existe_vertice(cancion) and grupos[cancion] == rdfy_constantes.GRUPO_CANCIONES:
        ciclo = funciones.encontrar_ciclo_largo_n(grafo, cancion, largo)
        if ciclo:
            print(" --> ".join(ciclo))
        else:
            print("No se encontro recorrido")
    else:
        print("El origen deberia ser una cancion")
        

##### COMANDO RANGO #####
def rango_canciones(grafo, distancia, cancion, grupos):
    if grafo.existe_vertice(cancion) and grupos[cancion] == rdfy_constantes.GRUPO_CANCIONES:
        print(funciones.encontrar_cantidad_a_distancia_n(grafo, cancion, distancia))
    else:
        print("El origen deberia ser una cancion")