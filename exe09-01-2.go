package main

import(
  "fmt"
  "sync"
  )
  
  var wg sync.WaitGroup
  
  func main(){
    gr:=100
    counter:=0
        wg.Add(gr)
    for i:=0; i < gr; i++{
  
      go fmt.Println("O valor da goroutine: ", i)
      counter+=1
      wg.Done()
    }
    wg.Wait()
    fmt.Println("contador: ", counter)
  }