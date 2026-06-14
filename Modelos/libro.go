package modelos

import "errors"

// Libro representa la estructura encapsulada de un libro.
type Libro struct {
	id        int
	titulo    string
	autor     string
	categoria string
}

// NuevoLibro es un constructor que valida los datos antes de crear la instancia (Encapsulación).
func NuevoLibro(id int, titulo, autor, categoria string) (Libro, error) {
	if id <= 0 {
		return Libro{}, errors.New("el ID debe ser un número entero positivo")
	}
	if titulo == "" || autor == "" {
		return Libro{}, errors.New("el título y el autor no pueden estar vacíos")
	}
	return Libro{
		id:        id,
		titulo:    titulo,
		autor:     autor,
		categoria: categoria,
	}, nil
}

// Métodos Getters para permitir la lectura segura de los campos encapsulados.

func (l Libro) ID() int           { return l.id }
func (l Libro) Titulo() string    { return l.titulo }
func (l Libro) Autor() string     { return l.autor }
func (l Libro) Categoria() string { return l.categoria }
