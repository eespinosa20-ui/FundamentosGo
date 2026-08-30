package main

import (
	"fmt"
	"math"
	"strconv"
	"time"
)

func main() {

	fmt.Println("--------------------------------------------------------------")
	fmt.Println("EJERCICIO 1:")
	const (
		DB_HOST       = "localhost"
		DB_PORT       = "5432"
		DB_MAX_CONN   = "25"
		DB_USE_SSL    = "true"
		DB_TIMEOUT_MS = "5000"
	)

	DB_HOST_VAR := DB_HOST
	fmt.Printf("DB_HOST: %v Tipo: %T\n", DB_HOST_VAR, DB_HOST_VAR)

	DB_PORT_VAR, err := strconv.Atoi(DB_PORT)
	if err != nil {
		fmt.Printf("Error al convertir DB_PORT: %v\n", err)
		DB_PORT_VAR = 5432
	}
	fmt.Printf("DB_PORT: %v Tipo: %T\n", DB_PORT_VAR, DB_PORT_VAR)

	DB_MAX_CONN_VAR, err := strconv.Atoi(DB_MAX_CONN)
	if err != nil {
		fmt.Printf("Error al convertir DB_MAX_CONN: %v\n", err)
		DB_MAX_CONN_VAR = 25
	}
	fmt.Printf("DB_MAX_CONN: %v Tipo: %T\n", DB_MAX_CONN_VAR, DB_MAX_CONN_VAR)

	DB_USE_SSL_VAR, err := strconv.ParseBool(DB_USE_SSL)
	if err != nil {
		fmt.Printf("Error al convertir DB_USE_SSL: %v\n", err)
		DB_USE_SSL_VAR = true
	}
	fmt.Printf("DB_USE_SSL: %v Tipo: %T\n", DB_USE_SSL_VAR, DB_USE_SSL_VAR)

	DB_TIMEOUT_MS_INT, err := strconv.Atoi(DB_TIMEOUT_MS)
	var DB_TIMEOUT_MS_VAR time.Duration
	if err != nil {
		fmt.Printf("Error parseando DB_TIMEOUT_MS: %v. Usando valor por defecto.\n", err)
		DB_TIMEOUT_MS_VAR = 3000 * time.Millisecond // Valor por defecto
	} else {
		DB_TIMEOUT_MS_VAR = time.Duration(DB_TIMEOUT_MS_INT) * time.Millisecond
	}
	fmt.Printf("DB_TIMEOUT_MS: %v Tipo: %T\n", DB_TIMEOUT_MS_VAR, DB_TIMEOUT_MS_VAR)

	fmt.Println("--------------------------------------------------------------")
	fmt.Println("EJERCICIO 2:")

	var num1 float64 = 19.99
	var num2 float64 = 5.50
	var num3 float64 = 12.00

	precio1 := int64(math.Round(num1 * 100))
	precio2 := int64(math.Round(num2 * 100))
	precio3 := int64(math.Round(num3 * 100))

	var producto1 int = 2
	var producto2 int = 1
	var producto3 int = 3

	totalINT := (precio1 * int64(producto1)) + (precio2 * int64(producto2)) + (precio3 * int64(producto3))
	totalFloat := float64(totalINT) / 100

	fmt.Printf("El total de la compra es: $%.2f\n", totalFloat)

	const IVA float64 = 0.16
	totalConIVA := totalFloat * (1 + IVA)
	fmt.Printf("El total de la compra con IVA es: $%.2f\n", totalConIVA)

	fmt.Println("--------------------------------------------------------------")
	fmt.Println("EJERCICIO 3:")

	var num int64 = 1734998400
	fecha := time.Unix(num, 0)

	fechaStr := fecha.Format("2006-01-02 15:04:05")
	var checkMark rune = '\u2705'

	fmt.Printf("%c Fecha: %v", checkMark, fechaStr)

}
