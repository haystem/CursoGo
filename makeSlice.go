package main

import ("fmt")

func main(){
  
  slice := make([]int, 5, 10) // make(tipo, len do slice, e capacidade)
  
  slice[1] = 2
  slice[2] = 3
  slice[3] = 5 
  slice[4] = 1
  
  fmt.Println(slice, len(slice),cap(slice))
  
  slice = append(slice, slice...)
  fmt.Println(slice, len(slice),cap(slice))
  slice = append(slice, slice...)
  fmt.Println(slice, len(slice),cap(slice))
}