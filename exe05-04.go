package main

import("fmt")

func main(){
  
  x:= struct{
    nome string
    endereco []string
  }{
    nome: "Alberto",
    endereco: []string{"Rua Manoel","Bairro: Guadalupe","Rio de Janeiro"},
  }
  
  fmt.Println(x)
}
