package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

// Connect inicializa de forma segura la conexión hacia la base de datos MySQL.
func Connect() (*sql.DB, error) {
	// Intentar cargar variables de entorno, si falla .env se maneja el error tolerando entornos de prod
	err := godotenv.Load()
	if err != nil {
		log.Println("Aviso: No se pudo cargar el archivo .env, se usarán variables del sistema.")
	}

	dns := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	// Crear la instancia de conexión con el driver de MySQL
	db, err := sql.Open("mysql", dns)
	if err != nil {
		return nil, fmt.Errorf("error al abrir conexión: %w", err)
	}

	// Verificar conectividad real mediante un Ping inmediato
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error de conectividad mediante ping: %w", err)
	}

	log.Println("Conexión a la base de datos establecida con éxito.")
	return db, nil
}
