package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

func main() {
    scanner := bufio.NewScanner(os.Stdin)

    fmt.Print("Introduce el primer número: ")
    scanner.Scan()
    input1 := strings.TrimSpace(scanner.Text())
    num1, err := strconv.ParseFloat(input1, 64)
    if err != nil {
        fmt.Println("Error: El valor introducido no es un número válido.")
        return
    }

    fmt.Print("Introduce el segundo número: ")
    scanner.Scan()
    input2 := strings.TrimSpace(scanner.Text())
    num2, err := strconv.ParseFloat(input2, 64)
    if err != nil {
        fmt.Println("Error: El valor introducido no es un número válido.")
        return
    }

    fmt.Print("Introduce la operación (+, -, *, /): ")
    scanner.Scan()
    operador := strings.TrimSpace(scanner.Text())

    var resultado float64

    switch operador {
    case "+":
        resultado = num1 + num2
    case "-":
        resultado = num1 - num2
    case "*":
        resultado = num1 * num2
    case "/":
        if num2 == 0 {
            fmt.Println("Error: No se puede dividir por cero.")
            return
        }
        resultado = num1 / num2
    default:
        fmt.Println("Error: Operador no válido.")
        return
    }

    fmt.Printf("Resultado: %.2f\n", resultado)
}
