package main

import ("fmt")

func main(){
  
  fmt.Println(fatorial(5))
  
}

func fatorial(x int) int{
   if x==1{
     return x
   }else{
    return x * fatorial(x-1)
   }
   
}