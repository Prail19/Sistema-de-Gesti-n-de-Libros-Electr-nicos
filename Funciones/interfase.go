package funciones

import "sistemalibros/modelos"

// RepositorioLibros define las operaciones permitidas para gestionar libros (Abstracción mediante Interfaces).
type RepositorioLibros interface {
	Agregar(libro modelos.Libro) error
	Listar() []modelos.Libro
	Buscar(titulo string) []modelos.Libro
	Eliminar(id int) error
}
