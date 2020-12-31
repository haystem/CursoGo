package main

import ("fmt")

func main(){
  x:= 2
  switch{
    case x< 5:
      fmt.Println("Xis é menor que 5")
    case x==5:
     fmt.Println("Xis é igual que 5")
    case x> 5:
     fmt.Println("Xis é maior que 5")
  }
}