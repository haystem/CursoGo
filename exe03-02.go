package main 

import ("fmt")

func main(){
  for n:=65; n <=90; n++{
    fmt.Println(n)
    fmt.Printf(" %s , %U\n", string(n),n)
    fmt.Printf(" %s , %U\n", string(n),n)
    fmt.Printf(" %s , %U\n", string(n),n)
    
  }
}