package main

import("fmt")

func main(){
  array := [5]int{1,2,3,4,5}
  
  for ind,valor := range array{
    fmt.Println(ind,valor)
  }
    fmt.Printf("%T",array)
}