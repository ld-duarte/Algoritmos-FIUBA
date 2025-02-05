from grafo import Grafo
import rdfy_constantes


def crear_grafo_gustos(datos):
    grafo = Grafo()
    grupo = {}
    for linea in datos:
        usuario = linea["USER_ID"]
        cancion = linea["TRACK_NAME"]
        artista = linea["ARTIST"]
        playlist = linea["PLAYLIST_NAME"]

        cancion_unica = f"{cancion} - {artista}"
        grafo.agregar_vertice(usuario)
        grafo.agregar_vertice(cancion_unica)
        grafo.agregar_arista(usuario, cancion_unica, playlist)
        
        grupo[usuario] = rdfy_constantes.GRUPO_USUARIOS
        grupo[cancion_unica] = rdfy_constantes.GRUPO_CANCIONES
        grupo[playlist] = rdfy_constantes.GRUPO_PLAYLISTS
    return grafo, grupo


def crear_grafo_canciones(datos):
    grafo_canciones = Grafo()
    usuario_canciones = {}
    for linea in datos:
        usuario = linea["USER_ID"]
        cancion = f'{linea["TRACK_NAME"]} - {linea["ARTIST"]}'
        if usuario not in usuario_canciones:
            usuario_canciones[usuario] = []
        usuario_canciones[usuario].append(cancion)
    for canciones in usuario_canciones.values():
        for i in range(len(canciones)):
            for j in range(i + 1, len(canciones)):
                grafo_canciones.agregar_vertice(canciones[i])
                grafo_canciones.agregar_vertice(canciones[j])
                grafo_canciones.agregar_arista(canciones[i], canciones[j])
    return grafo_canciones