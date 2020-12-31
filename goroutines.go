package main

import (
  "fmt"
  "sync"
  "time"
//  "runtime"
  )
var wg sync.WaitGroup
func main(){
  x:=100
  wg.Add(2)
  go soma1(x)
  go soma2(x)
  wg.Wait()  
  
}

func soma1(x int) {
    for i:=1; i < x; i++{
    fmt.Println("func1:",i)
    time.Sleep(1 + time.Second)
  }
  wg.Done()
}  

func soma2(x int) {
    for i:=0; i < x; i++{
    fmt.Println("func2",i+2)
    time.Sleep(1 + time.Second)
  }
  wg.Done()
}