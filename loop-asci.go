package main

import (
	"fmt"
)

func main() {
	for x := 33; x <= 122; x++ {
		fmt.Printf("O valor de %d em ASCI: %s \n", x, string(x)) //convertemos o número para string
	}
}