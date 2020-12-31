package main

import (
	"fmt"
)

type pessoa struct {
	nome      string
	sobrenome string
	hobbies   []string
}

func main() {
	adicionarValor("Carolina", "Ferraz", "Quitar", "Roubar kill", "jJungle")
	adicionarValor("Aarolina", "Ferraz", "Quitar", "Roubar kill", "jJungle")
}

func adicionarValor(x , y , k , w , z string) {
	var valor map[string]pessoa = make(map[string]pessoa)
	valor[y] = pessoa{
		nome:    x,
		hobbies: []string{k, w, z},
	}

	for chave, val := range valor {
		fmt.Println(chave, ":")
		fmt.Println("O hobbies preferidos:")
		for ind, num := range val.hobbies {
			fmt.Println("\t", ind+1, num)
		}
	}
}