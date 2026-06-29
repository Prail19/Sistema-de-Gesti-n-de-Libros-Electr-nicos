package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"sistemalibros/modelos"
)

type RespuestaAPI struct {
	Mensaje string      `json:"mensaje"`
	Data    interface{} `json:"data,omitempty"`
}

func escribirJSON(w http.ResponseWriter, estado int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(estado)
	json.NewEncoder(w).Encode(data)
}

// 1. GET /api/libros
func APIListarLibros(w http.ResponseWriter, r *http.Request) {
	libros := Repo.Listar()
	escribirJSON(w, http.StatusOK, convertirLibrosVista(libros))
}

// 2. GET /api/libros/{id}
func APIObtenerLibro(w http.ResponseWriter, r *http.Request) {
	idTexto := strings.TrimPrefix(r.URL.Path, "/api/libros/")
	id, err := strconv.Atoi(idTexto)
	if err != nil {
		escribirJSON(w, http.StatusBadRequest, RespuestaAPI{Mensaje: "ID inválido"})
		return
	}

	for _, libro := range Repo.Listar() {
		if libro.ID() == id {
			escribirJSON(w, http.StatusOK, LibroVista{
				ID:        libro.ID(),
				Titulo:    libro.Titulo(),
				Autor:     libro.Autor(),
				Categoria: libro.Categoria(),
			})
			return
		}
	}

	escribirJSON(w, http.StatusNotFound, RespuestaAPI{Mensaje: "Libro no encontrado"})
}

// 3. POST /api/libros
func APICrearLibro(w http.ResponseWriter, r *http.Request) {
	var libro LibroVista

	err := json.NewDecoder(r.Body).Decode(&libro)
	if err != nil {
		escribirJSON(w, http.StatusBadRequest, RespuestaAPI{Mensaje: "JSON inválido"})
		return
	}

	nuevoLibro, err := modelos.NuevoLibro(libro.ID, libro.Titulo, libro.Autor, libro.Categoria)
	if err != nil {
		escribirJSON(w, http.StatusBadRequest, RespuestaAPI{Mensaje: err.Error()})
		return
	}

	err = Repo.Agregar(nuevoLibro)
	if err != nil {
		escribirJSON(w, http.StatusBadRequest, RespuestaAPI{Mensaje: err.Error()})
		return
	}

	escribirJSON(w, http.StatusCreated, RespuestaAPI{
		Mensaje: "Libro creado correctamente",
		Data:    libro,
	})
}

// 4. PUT /api/libros/{id}
func APIActualizarLibro(w http.ResponseWriter, r *http.Request) {
	idTexto := strings.TrimPrefix(r.URL.Path, "/api/libros/")
	id, err := strconv.Atoi(idTexto)
	if err != nil {
		escribirJSON(w, http.StatusBadRequest, RespuestaAPI{Mensaje: "ID inválido"})
		return
	}

	var libroActualizado LibroVista
	err = json.NewDecoder(r.Body).Decode(&libroActualizado)
	if err != nil {
		escribirJSON(w, http.StatusBadRequest, RespuestaAPI{Mensaje: "JSON inválido"})
		return
	}

	_ = Repo.Eliminar(id)

	nuevoLibro, err := modelos.NuevoLibro(id, libroActualizado.Titulo, libroActualizado.Autor, libroActualizado.Categoria)
	if err != nil {
		escribirJSON(w, http.StatusBadRequest, RespuestaAPI{Mensaje: err.Error()})
		return
	}

	err = Repo.Agregar(nuevoLibro)
	if err != nil {
		escribirJSON(w, http.StatusBadRequest, RespuestaAPI{Mensaje: err.Error()})
		return
	}

	escribirJSON(w, http.StatusOK, RespuestaAPI{Mensaje: "Libro actualizado correctamente"})
}

// 5. DELETE /api/libros/{id}
func APIEliminarLibro(w http.ResponseWriter, r *http.Request) {
	idTexto := strings.TrimPrefix(r.URL.Path, "/api/libros/")
	id, err := strconv.Atoi(idTexto)
	if err != nil {
		escribirJSON(w, http.StatusBadRequest, RespuestaAPI{Mensaje: "ID inválido"})
		return
	}

	err = Repo.Eliminar(id)
	if err != nil {
		escribirJSON(w, http.StatusNotFound, RespuestaAPI{Mensaje: err.Error()})
		return
	}

	escribirJSON(w, http.StatusOK, RespuestaAPI{Mensaje: "Libro eliminado correctamente"})
}

// 6. GET /api/busqueda?q=texto
func APIBuscarLibros(w http.ResponseWriter, r *http.Request) {
	termino := r.URL.Query().Get("q")
	resultados := Repo.Buscar(termino)

	escribirJSON(w, http.StatusOK, convertirLibrosVista(resultados))
}

// 7. POST /api/registro
func APIRegistro(w http.ResponseWriter, r *http.Request) {
	escribirJSON(w, http.StatusOK, RespuestaAPI{
		Mensaje: "Servicio de registro disponible. Se puede conectar con MySQL usando RegistrarUsuario.",
	})
}

// 8. POST /api/login
func APILogin(w http.ResponseWriter, r *http.Request) {
	escribirJSON(w, http.StatusOK, RespuestaAPI{
		Mensaje: "Servicio de login disponible. Se puede conectar con MySQL usando IniciarSesion.",
	})
}

// Controlador principal para /api/libros y /api/libros/{id}
func APILibrosRouter(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/api/libros" {
		switch r.Method {
		case http.MethodGet:
			APIListarLibros(w, r)
		case http.MethodPost:
			APICrearLibro(w, r)
		default:
			escribirJSON(w, http.StatusMethodNotAllowed, RespuestaAPI{Mensaje: "Método no permitido"})
		}
		return
	}

	if strings.HasPrefix(r.URL.Path, "/api/libros/") {
		switch r.Method {
		case http.MethodGet:
			APIObtenerLibro(w, r)
		case http.MethodPut:
			APIActualizarLibro(w, r)
		case http.MethodDelete:
			APIEliminarLibro(w, r)
		default:
			escribirJSON(w, http.StatusMethodNotAllowed, RespuestaAPI{Mensaje: "Método no permitido"})
		}
		return
	}

	escribirJSON(w, http.StatusNotFound, RespuestaAPI{Mensaje: "Ruta no encontrada"})
}
