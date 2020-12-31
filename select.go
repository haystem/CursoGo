package main

import("fmt")

func main(){
   a:= make(chan int)
   b:= make(chan int)
   
   x:=500
   
   go func (x int){
     for i:=0; i < x; i++{
       a<-i
     }
   }(x/2)
   
  go func (x int){
     for i:=0; i < x; i++{
       b<-i
     }
   }(x/2)
   
   for j:=0; j < x; j++{
     select {
        case v:=<-a:
          fmt.Println("O Canal A:", v)
         case v:=<-b:
          fmt.Println("O Canal B:", v)
     }
   }
}