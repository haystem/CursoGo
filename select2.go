package main

import("fmt")

func main(){
  canal:= make(chan int)
  quit:= make(chan int)
  
  go receber(canal,quit)
  enviar(canal,quit)
  
  
}

func receber(c chan int, quit chan int){
  for i:=0; i < 20; i++{
     fmt.Println("O valor recebido: ", <-c)
  }
  quit <- 0
}

func enviar(c chan int, quit chan int){
   x:=0
   for{
     select{
       case c<- x:
        x++
       case <- quit:
     }
   }     return
}