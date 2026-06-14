package funciones

import (
	"fmt"
	"sistemalibros/modelos"
)

func MostrarLibros(libros []modelos.Libro) {
	for _, l := range libros {
		fmt.Println("ID:", l.ID)
		fmt.Println("Título:", l.Titulo)
		fmt.Println("Autor:", l.Autor)
		fmt.Println("Categoría:", l.Categoria)
		fmt.Println("----------------------")
	}
}
