package modelos

import "strings"

type Usuario struct {
	ID           int    `json:"id"`
	Nombre       string `json:"nombre"`
	Email        string `json:"email"`
	PasswordHash string `json:"password_hash"`
}

// Validar verifica si los campos esenciales del usuario cumplen las condiciones básicas.
func (u Usuario) Validar() bool {
	return u.Nombre != "" && strings.Contains(u.Email, "@")
}
