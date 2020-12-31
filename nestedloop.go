package main

import ("fmt")

func main (){
  for h :=0; h <=12;  h++{
   for m := 0; m <= 60 ; m++ {
        for s := 0; s <=60; s++ {
         fmt.Print(" ",s)
      }
       fmt.Print("\n\n",m)
    }
    fmt.Print("\n\n", h)
  } 
  fmt.Println(" ")
}