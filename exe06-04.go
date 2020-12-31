package main

import ("fmt")

type pessoa struct{
  nome string
  sobrenome string
  idade int
}

func (p pessoa) imprimirPessoa(){
  fmt.Printf("O meu nome %s, tenho sobrenome : %v, e minha idade %d", p.nome, p.sobrenome,p.idade)
}

func main(){
  
  pessoas := pessoa{
    nome: "Alberto",
    sobrenome : "Rocha",
    idade : 29, 
  }
  
  pessoas.imprimirPessoa()
}