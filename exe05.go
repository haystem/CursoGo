package main 

import (
  "fmt"
  )

//tipo criado
  type laranjas int8

//criado variavel global

  var x laranjas
  var y int8
  
  func main (){
    fmt.Println("O valor de X: ",x)
    fmt.Printf("O o tipo da variavel X: %T \n", x)
    x =42 
    fmt.Println("O valor de X: ",x)
    
    y = int8(x)
    fmt.Println("O valor de Y: ",y)
    fmt.Printf("O o tipo da variavel Y: %T \n", y)
    
  }