from collections import deque
import random

COEFICIENTE_AMORTIGUACION = 0.85
ITERACIONES_PR = 100
LARGO_CAMINO_ALEATORIO = 15
CANTIDAD_CAMINOS_ALEATORIOS = 5000


def camino_minimo_bfs(grafo, destino, origen):
    padre = {origen:None}
    cola = deque()
    cola.append(origen)

    while cola:
        actual = cola.popleft()
        if actual == destino:
            break
        
        for adyacente in grafo.obtener_adyacentes(actual):
            if adyacente not in padre:
                padre[adyacente] = actual
                cola.append(adyacente)
    return padre

def page_rank(grafo):
    n = len(grafo.obtener_vertices())
    padres = {}

    for v in grafo:
        for w in grafo.obtener_adyacentes(v):
            if w not in padres:
                padres[w] = []
            padres[w].append(v)
    
    pagerank = {v : 1/n for v in grafo}
    for _ in range(ITERACIONES_PR): 
        nuevo_pagerank = {}
        for v in grafo:
            nuevo_pagerank[v] = (1 - COEFICIENTE_AMORTIGUACION)/n
            if v in padres:
                sumatoria = 0
                for padre in padres[v]:
                    sumatoria += pagerank[padre] / len(grafo.obtener_adyacentes(padre))
                nuevo_pagerank[v] += COEFICIENTE_AMORTIGUACION * sumatoria
        pagerank = nuevo_pagerank
    return pagerank

def page_rank_personalizado(grafo, iniciales):
    pagerank = {nodo: 0 for nodo in grafo.obtener_vertices()}
    
    for _ in range(CANTIDAD_CAMINOS_ALEATORIOS):
        actual = random.choice(iniciales)
        valor = 1

        if not grafo.existe_vertice(actual):
            break

        for _ in range(LARGO_CAMINO_ALEATORIO):
            adyacentes = grafo.obtener_adyacentes(actual) 
            if not adyacentes:
                break

            siguiente = random.choice(adyacentes)
            valor /= len(adyacentes)
            pagerank[siguiente] += valor
            actual = siguiente
    
    for inicial in iniciales:
        pagerank.pop(inicial, None)

    return pagerank

def dfs_ciclo_largo_n(grafo, actual, origen, n, camino, visitados):
    if len(camino) == n:
        if origen in grafo.obtener_adyacentes(actual):
            camino.append(origen)
            return camino
        return None
    
    for vecino in grafo.obtener_adyacentes(actual):
        if vecino not in visitados:
            visitados.add(vecino)
            camino.append(vecino)
            ciclo = dfs_ciclo_largo_n(grafo, vecino, origen, n, camino, visitados)
            if ciclo:
                return ciclo
            camino.pop()
            visitados.remove(vecino)
    return None

def encontrar_ciclo_largo_n(grafo, origen, n):
    if n < 3:
        return None
    visitados = set()
    visitados.add(origen)
    camino = [origen]
    return dfs_ciclo_largo_n(grafo, origen, origen, n, camino, visitados)

def encontrar_cantidad_a_distancia_n(grafo, origen, n):
    distancia = {origen:0}
    contador = 0
    cola = deque()
    cola.append(origen)

    while cola:
        v = cola.popleft()
        for w in grafo.obtener_adyacentes(v):
            if w not in distancia:
                distancia[w] = distancia[v] + 1
                if distancia[w] == n:
                    contador += 1
                if distancia[w] < n:
                    cola.append(w)
    return contador

