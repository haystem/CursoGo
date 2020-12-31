package main

import("fmt")

type pessoa struct{
  nome string
  sobrenome string
  sabores []string
}

func main(){
  pessoa1:= pessoa{
      nome:"Alberto",
      sobrenome:"Rocha",
      sabores : []string{"Baunilha", "Chocolate"},
  }
  fmt.Println(pessoa1.nome)
  for _,x := range pessoa1.sabores{
    fmt.Println("\t",x)
  }
  
  pessoa2:= pessoa{
    nome:"Maria",
    sobrenome:"Rocha",
    sabores : []string{"Maracuja", "Morango"},
  }
  
  fmt.Println(pessoa2.nome)
  for _,x := range pessoa2.sabores{
    fmt.Println("\t",x)
  }
}

