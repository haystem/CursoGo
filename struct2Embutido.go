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
  
  fmt.Println(pessoa2)
  fmt.Println(pessoa1) 
}