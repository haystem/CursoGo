package main

import("fmt")

func main(){
 array := [10]string{"Ana","Joana","Siraya","Carla","Vivian","Betina","Laritssa","Bianca","Muana","Juressica"} 

  for indice, valor:= range array{
    fmt.Printf("%d, %s \n", indice, valor)
  }
  fmt.Printf("%T",array)
}
