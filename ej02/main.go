package main

import (
    "fmt"
)

func main() {
    var nombre string
    var edad int
    var altura float64

    fmt.Print("Introduce tu nombre, edad y altura (separados por espacios): ")

    // fmt.Scanln lee los valores en una sola línea
    _, err := fmt.Scanln(&nombre, &edad, &altura)
    if err != nil {
        fmt.Println("Error: Entrada no válida. Asegúrate de introducir los tres valores correctamente.")
        fmt.Println("Error devuelto por Go:", err)
        return
    }

    fmt.Println("\n--- Registro completado ---")
    fmt.Printf("Usuario: %s\n", nombre)
    fmt.Printf("Edad:    %d años\n", edad)
    fmt.Printf("Altura:  %.2fm\n", altura)
    fmt.Println("-------------------------")
}
