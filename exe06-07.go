package main

import ("fmt")

func main(){
  y:= func(x string) {
    fmt.Println("O valor chamado foi:", x)
  }
  
  y("solarizado")
}