class Grafo:
    
    def __init__(self, es_dirigido = False):
        self.vertices = {}
        self.es_dirigido = es_dirigido
    
    def agregar_vertice(self, vertice):
        if vertice not in self.vertices:
            self.vertices[vertice] = {}
    
    def agregar_arista(self, vertice_origen, vertice_destino, peso = 1):
        if not self.existe_vertice(vertice_origen) or not self.existe_vertice(vertice_destino):
            raise KeyError("los vertices deben existir en el grafo")
        self.vertices[vertice_origen][vertice_destino] = peso
        if not self.es_dirigido:
            self.vertices[vertice_destino][vertice_origen] = peso
    
    def existe_vertice(self, vertice):
        return vertice in self.vertices
    
    def existe_arista(self, vertice_origen, vertice_destino):
        return vertice_origen in self.vertices and vertice_destino in self.vertices[vertice_origen]
    
    def obtener_adyacentes(self, vertice):
        if not self.existe_vertice(vertice):
            raise KeyError(f"El vertice {vertice} no existe")
        return list(self.vertices[vertice].keys())

    def obtener_vertices(self):
        return list(self.vertices.keys())
    
    def eliminar_arista(self, vertice_origen, vertice_destino):
        if not self.existe_arista(vertice_origen, vertice_destino):
            raise KeyError(f"La arista {vertice_origen} --> {vertice_destino} no existe")  
        self.vertices[vertice_origen].pop(vertice_destino, None)
        if not self.es_dirigido:
            self.vertices[vertice_destino].pop(vertice_origen, None)
    
    def eliminar_vertice(self, vertice):
        if not self.existe_vertice(vertice):
            raise KeyError(f"El vertice {vertice} no existe")
        for adyacente in self.obtener_adyacentes(vertice):
            self.vertices[adyacente].pop(vertice, None)
        self.vertices.pop(vertice, None)

    def obtener_peso_arista(self, vertice_origen, vertice_destino):
        if not self.existe_arista(vertice_origen, vertice_destino):
            raise KeyError(f"La arista {vertice_origen} --> {vertice_destino} no existe")    
        return self.vertices[vertice_origen][vertice_destino]
        
    
    def __iter__(self):
        self._iter_vertices = iter(self.vertices.keys())
        return self
    
    def __next__(self):
        return next(self._iter_vertices)