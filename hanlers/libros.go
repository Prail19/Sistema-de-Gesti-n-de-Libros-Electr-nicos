package handlers

import (
	"html/template"
	"net/http"
	"strconv"

	"sistemalibros/funciones"
	"sistemalibros/modelos"
)

var Repo funciones.RepositorioLibros = funciones.NuevoGestorMemoria()

type LibroVista struct {
	ID        int
	Titulo    string
	Autor     string
	Categoria string
}

func convertirLibrosVista(libros []modelos.Libro) []LibroVista {
	var resultado []LibroVista

	for _, l := range libros {
		resultado = append(resultado, LibroVista{
			ID:        l.ID(),
			Titulo:    l.Titulo(),
			Autor:     l.Autor(),
			Categoria: l.Categoria(),
		})
	}

	return resultado
}

func RenderTemplate(w http.ResponseWriter, archivo string, data interface{}) {
	tmpl, err := template.ParseFiles(
		"templates/base.html",
		"templates/"+archivo,
	)

	if err != nil {
		http.Error(w, "Error al cargar template: "+err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.ExecuteTemplate(w, "base", data)
	if err != nil {
		http.Error(w, "Error al ejecutar template: "+err.Error(), http.StatusInternalServerError)
	}
}

func InicioHandler(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "inicio.html", nil)
}

func LibrosHandler(w http.ResponseWriter, r *http.Request) {
	libros := Repo.Listar()
	RenderTemplate(w, "libros.html", convertirLibrosVista(libros))
}

func AgregarHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		RenderTemplate(w, "agregar.html", nil)
		return
	}

	id, err := strconv.Atoi(r.FormValue("id"))
	if err != nil {
		RenderTemplate(w, "agregar.html", "ID inválido")
		return
	}

	titulo := r.FormValue("titulo")
	autor := r.FormValue("autor")
	categoria := r.FormValue("categoria")

	libro, err := modelos.NuevoLibro(id, titulo, autor, categoria)
	if err != nil {
		RenderTemplate(w, "agregar.html", err.Error())
		return
	}

	err = Repo.Agregar(libro)
	if err != nil {
		RenderTemplate(w, "agregar.html", err.Error())
		return
	}

	http.Redirect(w, r, "/libros", http.StatusSeeOther)
}

func BuscarHandler(w http.ResponseWriter, r *http.Request) {
	termino := r.URL.Query().Get("q")
	libros := Repo.Buscar(termino)

	RenderTemplate(w, "libros.html", convertirLibrosVista(libros))
}

func EliminarHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err == nil {
		_ = Repo.Eliminar(id)
	}

	http.Redirect(w, r, "/libros", http.StatusSeeOther)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		RenderTemplate(w, "login.html", nil)
		return
	}

	http.Redirect(w, r, "/libros", http.StatusSeeOther)
}

func RegistroHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		RenderTemplate(w, "registro.html", nil)
		return
	}

	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
