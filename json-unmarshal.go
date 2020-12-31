package main

import (
  "fmt"
  "encoding/json"
  //"os"
  )
  
    func main (){
    json1:= []byte(`{"ID":1,"Name":"Quentes","Colors":["Vemelho","Laranja","Rosa","Vinho"]}`)
 
    type informacao struct{
      ID int 
      Name string
      Colors []string 
    } 
    
    var conver informacao
    
    err := json.Unmarshal(json1, &conver)
  	if err != nil {
		fmt.Println("error:", err)
    	}
    	fmt.Printf("%+v \n", conver)
      fmt.Printf("%+v \n",conver.ID)
  }