package main

import (
  "fmt"
)
  func main(){
    x := "é"
    y := "e"
    z := "uç99"
    
    fmt.Println(z , y , x)
    
    a := []byte (x)
    b := []byte (y)
    c := []byte (z)
    
    fmt.Println(a,b,c)
}