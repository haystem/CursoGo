package main

import("fmt")

func main(){
  
  vmp := map[int]string{
      1:"Mirella",
      2:"Fernanda",
      3:"Ruth",
      4:"Raquel",
      5:"Mirelia",
  }
  
  fmt.Println(vmp)
  fmt.Println(vmp[1])
  
  if vmp,ok := vmp[7]; !ok{
    fmt.Println("Não possui",vmp)
    
  }else{
    fmt.Println("Possui", vmp)
  }
  
}