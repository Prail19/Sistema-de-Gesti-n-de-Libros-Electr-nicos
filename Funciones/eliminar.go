package funciones

import "sistemalibros/modelos"

func EliminarLibro(libros []modelos.Libro, id int) []modelos.Libro {
	var nuevos []modelos.Libro

	for _, l := range libros {
		// CORRECCIÓN: Agregamos () después de ID
		if l.ID() != id {
			nuevos = append(nuevos, l)
		}
	}
	return nuevos
}
