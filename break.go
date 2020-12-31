package main 

import("fmt")

func main (){
  x:=0
  for {
    if x < 10{
      x++
      fmt.Println("Xis menor que 10")
    } else {
      fmt.Println("Xis maior que 10, PAROUU")
      break
    }
  }
}