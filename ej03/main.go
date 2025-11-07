package main

import (
    "flag"
    "fmt"
    "os"
)

func main() {
    // Definir los flags
    fileName := flag.String("f", "output.txt", "Nombre del fichero a crear")
    content := flag.String("c", "Hello, Go!", "Contenido a escribir en el fichero")
    repeat := flag.Int("n", 1, "Número de veces que se escribirá el contenido")

    // Procesar los flags
    flag.Parse()

    // Crear o sobrescribir el fichero
    file, err := os.Create(*fileName)
    if err != nil {
        fmt.Println("Error al crear el fichero:", err)
        return
    }
    defer file.Close()

    // Escribir el contenido n veces
    for i := 0; i < *repeat; i++ {
        _, err := file.WriteString(fmt.Sprintf("%s\n", *content))
        if err != nil {
            fmt.Println("Error al escribir en el fichero:", err)
            return
        }
    }

    fmt.Printf("Fichero '%s' generado correctamente con %d línea(s).\n", *fileName, *repeat)
}
