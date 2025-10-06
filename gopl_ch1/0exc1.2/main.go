// Modifique o programa echo para exibir o indice e o valor de cada um de seus argumentos, linha por linha
package main

import (
	"fmt"
	"os"
)

func main() {
	for i, arg := range os.Args {
		fmt.Println("Indice: ", i, "Caractere: ", arg)
	}
}
