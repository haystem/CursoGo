package main

import(
  "fmt"
  )

func main(){
  
  canal := make(chan int)
  funcao := 100
  
  go func(c chan<- int){
    for i:=0 ; ; i++{
       c <- i+1
    }
    close(c)
  }(canal)
  
  imprimirValor(funcao, canal)
  
}

func imprimirValor(x int, c <-chan int){
   for i:=0; i < x; i++{
     fmt.Println("O valor do canal: ", <-c)
   }
}