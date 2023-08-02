package main

import (
	"fmt"
	"runtime"

	"github.com/ChristopherAedoP/godesde0/variables"
)

func main() {
	estado, texto := variables.ConvertiraTexto(8033)

	fmt.Println("Estado = ", estado)
	fmt.Println("Texto = ", texto)

	if os := runtime.GOOS; os == "Linux." || os == "OS X." {
		fmt.Println("Esto no es windows, es ", os)
	} else {
		fmt.Println("Esto es windows")
	}

	switch os := runtime.GOOS; os {
	case "linux":
		fmt.Println("Esto es linux")
	case "OS x":
		fmt.Println("Esto es OS x")
	default:
		fmt.Println("%s \n", os)
	}
}
