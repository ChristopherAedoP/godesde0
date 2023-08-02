package main

import (
	"fmt"

	"github.com/ChristopherAedoP/godesde0/variables"
)

func main() {
	estado, texto := variables.ConvertiraTexto(8033)

	fmt.Println("Estado = ", estado)
	fmt.Println("Texto = ", texto)
}
