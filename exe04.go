package main 

import (
  "fmt"
  )

  type laranjas int8
  var x laranjas
  func main (){
    x = 42
    
    fmt.Printf("O valor de X: %v do tipo %T.",x,x)
   
  }