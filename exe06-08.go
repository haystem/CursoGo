package main 

import ("fmt")

func main(){
  y := receber()

  fmt.Println(y(5))  
}

func receber() func(x int) int{
  return func(x int) int{
    return x*(x+2)

  }
  }