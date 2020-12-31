package main

import(
  "fmt"
  "time"
  )

func main(){
  canal:= converge(trabalho("mara"),trabalho("moro"))
  
  for i:=0; i < 10; i++{
    fmt.Println(<-canal)
  }
  
  
}

func trabalho (s string) chan string{
  canal := make(chan string)
  go func(a string, can chan string){
    for i:=0; ; i++{
       can <- fmt.Sprintf("Os valores da %s : %v",s,i)
       time.Sleep(1 * time.Second)
    }
  }(s,canal)
       fmt.Print(".")
       time.Sleep(1 * time.Second)
       fmt.Print("..")
       time.Sleep(1 * time.Second)
       fmt.Print("... \n")
       time.Sleep(1 * time.Second)
  
  return canal
}
func converge(x, y chan string) chan string{
   t := make(chan string)
   go func(){
     for{
       t <- <- x 
     }
   }()
   
   go func(){
    for{
      t <- <- y 
    }
   }()
   return t
}