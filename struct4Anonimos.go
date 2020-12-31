package main

import("fmt")

func main(){
  x := struct{
    nome string
    idade int
  }{
    nome:"Alberto",
    idade:29,
  }
  fmt.Println(x)
}