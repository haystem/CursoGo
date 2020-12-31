package main 

import (
      "fmt"
      "encoding/json"
      "os"
  )
  
type colorGroup struct{
  ID int
  Name string
  Colors []string
}  
  
  func main(){
    
    group := colorGroup{
      ID :1,
      Name : "Quentes",
      Colors : []string{"Vemelho","Laranja","Rosa","Vinho"},
    }
    
    b,err := json.Marshal(group)
    if err != nil{
      fmt.Println("error",err)
    }
    os.Stdout.Write(b)
  }