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

  pessoa2:= pessoa{
    nome:"Maria",
    sobrenome:"Rocha",
    sabores : []string{"Maracuja", "Morango"},
  }
  
  fmt.Println(pessoa1.nome)
  imprimirSabores(pessoa1.sabores)
  fmt.Println(pessoa2.nome) 
  imprimirSabores(pessoa2.sabores)
 
}

func imprimirSabores(condicao []string){
  for _,y:= range condicao{
  fmt.Println("\t",y)
  }
}
