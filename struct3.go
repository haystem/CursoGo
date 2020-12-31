package main

import("fmt")

type pessoa struct{
  nome string
  idade int
}

type profissional struct{
  pessoa  // está se referindo a type pessoa anterior caso queira mesclar
  título string
  salario float64 //sempre indicar float64 
}
func main(){
  pessoa1:=pessoa{
    nome:"Alfredo",
    idade:30,
  }
  
  pessoa2:= profissional{
    pessoa : pessoa{  // informei que pessoa recebe os campos de struct pessoa
      nome:"Maria",
      idade:28,
    },
    título:"Enfermeira",
    salario:2800,
  }
  
  pessoa3:=pessoa{"Mauricio",34}
  pessoa4:=profissional{pessoa{"Ana Maria",38},"Politica",28374847}
  
  fmt.Println(pessoa2.idade,pessoa2.pessoa.nome)
  fmt.Println(pessoa1) 
  fmt.Println(pessoa3) 
  fmt.Println(pessoa4) 
  fmt.Println("")
  fmt.Printf("\n %T \n",pessoa4)   
}