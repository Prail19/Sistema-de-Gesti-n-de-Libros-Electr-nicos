package main

import (
	"log"
	"net/http"

	"sistemalibros/db"
	"sistemalibros/handlers"
)

func main() {
	database, err := db.Connect()
	if err != nil {
		log.Fatal("Error al conectar con la base de datos: ", err)
	}
	defer database.Close()

	// Archivos estáticos: imágenes, CSS, JS
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Rutas de la aplicación web
	http.HandleFunc("/", handlers.InicioHandler)
	http.HandleFunc("/libros", handlers.LibrosHandler)
	http.HandleFunc("/agregar", handlers.AgregarHandler)
	http.HandleFunc("/buscar", handlers.BuscarHandler)
	http.HandleFunc("/eliminar", handlers.EliminarHandler)
	http.HandleFunc("/login", handlers.LoginHandler)
	http.HandleFunc("/registro", handlers.RegistroHandler)

	// Servicios Web API JSON
	http.HandleFunc("/api/libros", handlers.APILibrosRouter)
	http.HandleFunc("/api/libros/", handlers.APILibrosRouter)
	http.HandleFunc("/api/busqueda", handlers.APIBuscarLibros)
	http.HandleFunc("/api/login", handlers.APILogin)
	http.HandleFunc("/api/registro", handlers.APIRegistro)

	log.Println("====================================")
	log.Println("Servidor iniciado correctamente")
	log.Println("http://localhost:3000")
	log.Println("====================================")

	log.Fatal(http.ListenAndServe(":3000", nil))
}
