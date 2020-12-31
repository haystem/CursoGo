package main

import("fmt")

func main(){
  
  x:= retornarFuncao()
  
  fmt.Println(x(3))
  fmt.Println(retornarFuncao()(4))
  
}

func retornarFuncao() func(int) int{
  return func(x int) int{
    return x*98398
  }
}