package main

import("fmt")

func main(){
  canal := make(chan int)
  
  go meuloop(10, canal)
  
  for v := range canal{
    fmt.Println("O valor do canal: ",v)
  }
  
}

func meuloop(s int, c chan<- int){
  for i:=0; i < s; i++{
     c <- i
  }
  close(c)
}  
