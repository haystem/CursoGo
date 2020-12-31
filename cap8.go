package main

import("fmt")

func main (){
  slice := []int{9,8,4,3,5,6,12,3,4}

  for x:=0; x < len(slice); x++{
    fmt.Println(slice[x])
    
    //uma forma de excluir item
    
    slice = append(slice[:3], slice[4:]...)
    fmt.Println(slice)
  }
}