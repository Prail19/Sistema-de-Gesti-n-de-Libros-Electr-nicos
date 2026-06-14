package funciones

import (
	"errors"
	"sistemalibros/modelos"
)

func AgregarLibro(libros []modelos.Libro, libro modelos.Libro) ([]modelos.Libro, error) {
	for _, l := range libros {
		// CORRECCIÓN: Asegurar el uso de () si se valida el ID encapsulado
		if l.ID() == libro.ID() {
			return libros, errors.New("el ID ya existe")
		}
	}
	return append(libros, libro), nil
}
