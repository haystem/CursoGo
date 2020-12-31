package main

import (
  "fmt"
  "os"
  "log"
  )
  
func main(){
  
  f, err := os.Create("log.txt")
  if err != nil{
    fmt.Println("error", err)
  } 
  
  defer f.Close()
  
  log.SetOutput(f)
  
 
  f2, err:= os.Open("mecanica.txt")
  if err != nil{
    log.Println("O erro", err)
  } 
  
  defer f2.Close()
  
  fmt.Println("Checar arquivo log")
}