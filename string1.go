package main

import (
  "fmt"
  )
 const sample = "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98"
func main(){
  X:= "Sistema"
  fmt.Println(X)
  X = "OLA"
  fmt.Println(X)
  
 fmt.Println(sample)
 for i := 0; i < len(sample); i++ {
   fmt.Printf("%b ", sample[i])
   }
  
}