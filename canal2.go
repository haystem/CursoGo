package main

import (
  "fmt"
  "sync"
  
  )
var wg sync.WaitGroup

func main(){
  wg.Add(2)
  defer wg.Wait()  
  canal := make(chan int)
  
    go enviar(canal)
    go receber(canal)
  
}

func enviar(s chan<- int) {
    s <- 555
    wg.Done()
}

func receber(r <-chan int){
  fmt.Println(<-r)
  wg.Done()
}