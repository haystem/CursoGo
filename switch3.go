package main

import ("fmt")

func main(){
  persona:= "jorge"
  switch persona{
    case "jorge","miranda":
      fmt.Println("Hoje que é responsável é jorge")
      fmt.Println("Hoje tbm quem é responsável é miranda")
    case "maria", "ana":
      fmt.Println("Hoje que é responsável é maria")
      fmt.Println("Hoje tbm quem é responsável é ana")
  }
}