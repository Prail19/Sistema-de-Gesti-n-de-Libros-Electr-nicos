package funciones

import (
	"errors"
	"sistemalibros/modelos"
	"strings"
)

// GestorMemoria implementa la interfaz RepositorioLibros usando un slice.
type GestorMemoria struct {
	libros []modelos.Libro
}

// NuevoGestorMemoria inicializa el repositorio.
func NuevoGestorMemoria() *GestorMemoria {
	return &GestorMemoria{libros: make([]modelos.Libro, 0)}
}

// Agregar inserta un libro controlando que el ID no se encuentre repetido (Manejo de Errores).
func (g *GestorMemoria) Agregar(libro modelos.Libro) error {
	for _, l := range g.libros {
		if l.ID() == libro.ID() {
			return errors.New("error: ya existe un libro registrado con este ID")
		}
	}
	g.libros = append(g.libros, libro)
	return nil
}

// Listar retorna todos los elementos almacenados.
func (g *GestorMemoria) Listar() []modelos.Libro {
	return g.libros
}

// Buscar filtra y localiza coincidencias ignorando diferencias entre mayúsculas y minúsculas.
func (g *GestorMemoria) Buscar(titulo string) []modelos.Libro {
	var resultados []modelos.Libro
	for _, l := range g.libros {
		if strings.Contains(strings.ToLower(l.Titulo()), strings.ToLower(titulo)) {
			resultados = append(resultados, l)
		}
	}
	return resultados
}

// Eliminar remueve un elemento por ID. Si no existe, retorna un error descriptivo.
func (g *GestorMemoria) Eliminar(id int) error {
	for i, l := range g.libros {
		if l.ID() == id {
			// Eliminación eficiente del elemento reestructurando el slice
			g.libros = append(g.libros[:i], g.libros[i+1:]...)
			return nil
		}
	}
	return errors.New("error: no se encontró ningún libro con el ID especificado")
}
