package main

import("fmt")

func main(){
  defer fmt.Println(1) // executa por ultimo
  fmt.Println(2)
  defer fmt.Println(3) //executa por ultimo
  fmt.Println(4)
  fmt.Println(5)
  fmt.Println(6)
}