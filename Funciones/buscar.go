package funciones

import (
	"sistemalibros/modelos"
	"strings"
)

func BuscarLibroPorTitulo(libros []modelos.Libro, titulo string) []modelos.Libro {
	var resultados []modelos.Libro

	for _, l := range libros {
		// CORRECCIÓN: Agregamos () después de Titulo
		if strings.Contains(strings.ToLower(l.Titulo()), strings.ToLower(titulo)) {
			resultados = append(resultados, l)
		}
	}
	return resultados
}
