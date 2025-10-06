// Modifique o program echo para exibir também o nome do programa
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Executando: ", strings.Join(os.Args[0:1], " "), strings.Join(os.Args[1:], " "))
}
