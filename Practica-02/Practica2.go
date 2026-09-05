package main

import (
	"fmt"
	"strings"
)

func main() {
	println("Ejercicio 1: Registro de usuarios")

	var username string = "ana"
	email := "ana@gmail"
	password := "1234"
	age := 17

	var errors []string

	if len(username) < 4 || len(username) > 20 {
		errors = append(errors, "Username invalido")
	}

	if !strings.Contains(email, "@") || !strings.Contains(email, ".") {
		errors = append(errors, "Email invalido")
	}

	if len(password) < 8 {
		errors = append(errors, "Password invalido")
	}

	if age <= 18 {
		errors = append(errors, "Edad invalida")
	}

	if len(errors) == 0 {
		fmt.Println("✅ Registro exitoso")
	} else {
		fmt.Println("❌ Error: ")
		fmt.Println(errors)
	}

	println("----------------------------------------------------")
	println("Ejercicio 2: HTTP Status Code")

	status := 418

	switch status / 100 {
	case 2:
		fmt.Println("Exito")
	case 3:
		fmt.Println("Redirección")
	case 4:
		fmt.Println("Error del cliente")
	case 5:
		fmt.Println("Error del servidor")
	default:
		fmt.Println("Código desconocido")
	}

}
