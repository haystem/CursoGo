package main

import ("fmt")

func main(){
  
  s1 := func () int{
      return 2+2
  }
  
  fmt.Println(s1)
}