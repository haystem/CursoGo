package main

import (
  "fmt"
  "sync"
  "runtime"
  "time"
  )
  
  var wg sync.WaitGroup
  
  func main(){
    wg.Add(4)
    
    go imprimirValor2(false)
    go imprimirValor3(5.675)
    go imprimirValor1(5)
    go imprimirValor4("hello")
    
    wg.Wait()
    
  }
  
  func imprimirValor1(x int) {
    fmt.Println("O Valor inserido foi: ",x)
    time.Sleep(3 * time.Second)
    wg.Done()
  }
  
    func imprimirValor2(x bool) {
    fmt.Println("O Valor inserido foi: ",x)
    runtime.Gosched()
    wg.Done()
  }
  
    func imprimirValor3(x float64) {
    fmt.Println("O Valor inserido foi: ",x)
    runtime.Gosched()
    wg.Done()
  }
  
    func imprimirValor4(x string) {
    fmt.Println("O Valor inserido foi: ",x)
    time.Sleep(8 * time.Second)
    wg.Done()
  }