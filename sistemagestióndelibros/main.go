package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"sistemalibros/db"
	"sistemalibros/funciones"
	"sistemalibros/modelos"
)

func main() {
	// Inicializar conexión segura a la Base de Datos usando nuestro paquete db
	database, err := db.Connect()
	if err != nil {
		log.Fatal("❌ Error crítico del sistema al conectar a la BD: ", err)
	}
	defer database.Close()

	// Inicializar el repositorio utilizando el contrato de la Interfaz
	var repo funciones.RepositorioLibros = funciones.NuevoGestorMemoria()
	reader := bufio.NewReader(os.Stdin)

	var usuarioLogueado modelos.Usuario
	autenticado := false

	// ==========================================
	// 1. FLUJO DE AUTENTICACIÓN SECURE LOG
	// ==========================================
	for !autenticado {
		fmt.Println("\n===== ACCESO AL SISTEMA =====")
		fmt.Println("1. Iniciar Sesión")
		fmt.Println("2. Registrarse")
		fmt.Println("3. Salir")
		fmt.Print("Seleccione una opción: ")

		opcionAuthTexto, _ := reader.ReadString('\n')
		opcionAuthTexto = strings.TrimSpace(opcionAuthTexto)
		opcionAuth, _ := strconv.Atoi(opcionAuthTexto)

		switch opcionAuth {
		case 1:
			fmt.Print("Ingrese Email: ")
			email, _ := reader.ReadString('\n')
			email = strings.TrimSpace(email)

			fmt.Print("Ingrese Contraseña: ")
			password, _ := reader.ReadString('\n')
			password = strings.TrimSpace(password)

			usuario, err := funciones.IniciarSesion(database, email, password)
			if err != nil {
				fmt.Printf("❌ Error: %v\n", err)
			} else {
				usuarioLogueado = usuario
				autenticado = true
				fmt.Printf("\n✨ ¡Bienvenido/a %s! Conexión autorizada.\n", usuarioLogueado.Nombre)
			}

		case 2:
			var nuevoUsuario modelos.Usuario
			fmt.Print("Ingrese Nombre Completo: ")
			nuevoUsuario.Nombre, _ = reader.ReadString('\n')
			nuevoUsuario.Nombre = strings.TrimSpace(nuevoUsuario.Nombre)

			fmt.Print("Ingrese Email Corporativo: ")
			nuevoUsuario.Email, _ = reader.ReadString('\n')
			nuevoUsuario.Email = strings.TrimSpace(nuevoUsuario.Email)

			fmt.Print("Ingrese Contraseña: ")
			passwordPlano, _ := reader.ReadString('\n')
			passwordPlano = strings.TrimSpace(passwordPlano)

			err := funciones.RegistrarUsuario(database, nuevoUsuario, passwordPlano)
			if err != nil {
				fmt.Printf("❌ Error en registro: %v\n", err)
			} else {
				fmt.Println("✨ Registro exitoso en MySQL. Proceda a iniciar sesión.")
			}

		case 3:
			fmt.Println("Saliendo de la aplicación.")
			return
		default:
			fmt.Println("Opción no válida.")
		}
	}

	// ==========================================
	// 2. MENÚ PRINCIPAL (USANDO INTERFACES Y ENCAPSULACIÓN)
	// ==========================================
	for {
		fmt.Println("\n===== GESTIÓN DE BIBLIOTECA (U3) =====")
		fmt.Println("1. Agregar nuevo libro")
		fmt.Println("2. Mostrar catálogo completo")
		fmt.Println("3. Buscar libro por título")
		fmt.Println("4. Eliminar libro del sistema")
		fmt.Println("5. Salir")
		fmt.Print("Seleccione una opción: ")

		opcionTexto, _ := reader.ReadString('\n')
		opcionTexto = strings.TrimSpace(opcionTexto)
		opcion, _ := strconv.Atoi(opcionTexto)

		switch opcion {
		case 1:
			fmt.Print("Ingrese ID Numérico: ")
			idTexto, _ := reader.ReadString('\n')
			idTexto = strings.TrimSpace(idTexto)
			id, _ := strconv.Atoi(idTexto)

			fmt.Print("Ingrese Título: ")
			titulo, _ := reader.ReadString('\n')
			titulo = strings.TrimSpace(titulo)

			fmt.Print("Ingrese Autor: ")
			autor, _ := reader.ReadString('\n')
			autor = strings.TrimSpace(autor)

			fmt.Print("Ingrese Categoría: ")
			categoria, _ := reader.ReadString('\n')
			categoria = strings.TrimSpace(categoria)

			// Aplicación de Encapsulación mediante el Constructor Seguro
			nuevoLibro, err := modelos.NuevoLibro(id, titulo, autor, categoria)
			if err != nil {
				fmt.Printf("❌ Error de Validación: %v\n", err)
				continue
			}

			// Inserción controlada mediante el contrato de la interfaz
			err = repo.Agregar(nuevoLibro)
			if err != nil {
				fmt.Printf("❌ Error de Almacenamiento: %v\n", err)
			} else {
				fmt.Println("✨ Libro catalogado con éxito.")
			}

		case 2:
			libros := repo.Listar()
			if len(libros) == 0 {
				fmt.Println("El catálogo se encuentra vacío.")
				continue
			}
			fmt.Println("\n--- Catálogo Disponible ---")
			for _, l := range libros {
				// Acceso mediante métodos de lectura encapsulados (Getters)
				fmt.Printf("ID: %d | Título: %s | Autor: %s | Categoría: %s\n", l.ID(), l.Titulo(), l.Autor(), l.Categoria())
			}

		case 3:
			fmt.Print("Escriba el título o término de búsqueda: ")
			busqueda, _ := reader.ReadString('\n')
			busqueda = strings.TrimSpace(busqueda)

			resultados := repo.Buscar(busqueda)
			if len(resultados) == 0 {
				fmt.Println("No se hallaron coincidencias.")
				continue
			}
			for _, l := range resultados {
				fmt.Printf("🔍 ID: %d - %s (%s)\n", l.ID(), l.Titulo(), l.Categoria())
			}

		case 4:
			fmt.Print("ID del libro a dar de baja: ")
			idTexto, _ := reader.ReadString('\n')
			idTexto = strings.TrimSpace(idTexto)
			id, _ := strconv.Atoi(idTexto)

			err := repo.Eliminar(id)
			if err != nil {
				fmt.Printf("❌ %v\n", err)
			} else {
				fmt.Println("✨ El libro ha sido removido del registro.")
			}

		case 5:
			fmt.Printf("Sesión de %s cerrada. ¡Adiós!\n", usuarioLogueado.Nombre)
			return
		default:
			fmt.Println("Opción inválida.")
		}
	}
}
