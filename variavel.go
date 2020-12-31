package main 

import (
    "fmt"
  )
  
  var y = 10  //package level scope
  func main (){
    z := 444
    fmt.Println(y)
    valor(z)
  }
  
  func valor(x int){
    fmt.Println(x)
    fmt.Println(y)
  }