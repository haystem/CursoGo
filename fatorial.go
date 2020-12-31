package main

import ("fmt")

func main(){
  
  fmt.Println(fatorial(5))
  
}

func fatorial(x int) int{
  fat := x
  
  for i := 1; i <= (x-1); i++{
     fat *= i
  }
  return fat
}