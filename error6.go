package main

import (
  "fmt"
  "log"
  )
  
  func main(){
    _, error := sqr(-120)
    if error != nil{
      log.Fatalln(error)
    }
  }
  
  func sqr(f int) (int, error){
    if f <= 0 {
      return 0, fmt.Errorf("Deu ruim com o número %v", f)
    }
    
    return 42, nil
  }