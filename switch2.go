package main

import ("fmt")

func main(){
  persona:= "jorge"
  switch persona{
    case "jorge":
      fmt.Println("Hoje que é responsável é jorge")
      fallthrough // pula a comparação e executsa o comando abaixo
    case "miranda":
     fmt.Println("Hoje que é responsável é miranda")
    case "maria":
      fmt.Println("Hoje que é responsável é maria")
    case "ana":
     fmt.Println("Hoje que é responsável é ana")
  }
}