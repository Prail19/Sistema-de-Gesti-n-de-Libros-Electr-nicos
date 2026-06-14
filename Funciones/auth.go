package funciones

import (
	"database/sql"
	"errors"
	"fmt"
	"sistemalibros/modelos"

	"golang.org/x/crypto/bcrypt"
)

// RegistrarUsuario valida los campos del struct y encripta el password mediante bcrypt.
func RegistrarUsuario(db *sql.DB, usuario modelos.Usuario, passwordPlano string) error {
	if !usuario.Validar() {
		return errors.New("datos de usuario inválidos, verifique el formato del correo")
	}
	if len(passwordPlano) < 6 {
		return errors.New("la contraseña debe poseer al menos 6 caracteres")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(passwordPlano), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("error al procesar la seguridad de la contraseña: %w", err)
	}

	_, err = db.Exec("INSERT INTO usuarios (nombre, email, password_hash) VALUES (?, ?, ?)",
		usuario.Nombre, usuario.Email, string(hashedPassword))
	if err != nil {
		return fmt.Errorf("error al insertar registro en la BD: %w", err)
	}
	return nil
}

// IniciarSesion comprueba credenciales en la base de datos de manera segura.
func IniciarSesion(db *sql.DB, email, passwordPlano string) (modelos.Usuario, error) {
	var usuario modelos.Usuario
	var hashedPassword string

	err := db.QueryRow("SELECT id, nombre, email, password_hash FROM usuarios WHERE email = ?", email).
		Scan(&usuario.ID, &usuario.Nombre, &usuario.Email, &hashedPassword)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return modelos.Usuario{}, errors.New("credenciales inválidas: el correo electrónico no está registrado")
		}
		return modelos.Usuario{}, fmt.Errorf("error en la consulta del sistema: %w", err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(passwordPlano))
	if err != nil {
		return modelos.Usuario{}, errors.New("credenciales inválidas: contraseña incorrecta")
	}

	return usuario, nil
}
