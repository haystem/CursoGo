package main 

import(
  "fmt"
  "sync"
  "time"
  )
  var wg sync.WaitGroup
  var mu sync.Mutex
func main(){
  contador :=1
  totalgoroutines :=2

  wg.Add(totalgoroutines)
  
  go soma(contador)
  go soma2(contador)
  
  fmt.Println(contador)
  wg.Wait()
}

func soma (x int) int {
    mu.Lock()
    a:= 0
    for b:=0; b < 10; b++ {
      a+=b
      time.Sleep(20)
      wg.Done()
    }
    mu.Unlock()
   return a
}

func soma2 (x int) int {
    mu.Lock()
    a:= 0
    for b:=0; b < 10; b++ {
      a+=b
      time.Sleep(20)
      wg.Done()
    } 
    mu.Unlock()
    return a
}
