package main

import (
  "fmt"
  )
  
  func main (){
    x:= "ola"
    fmt.Print(x) // print na tela sem dar linha
    
    y:= "Bom dia"
    fmt.Println(y) //entre cada item tem um espaço
    
    z:= fmt.Sprint(x,y)  // retorna tipo string  o valor vai ser salvo em string
    fmt.Println(z)
  }