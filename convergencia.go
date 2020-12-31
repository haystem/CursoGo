package main

import(
  "fmt"
 // "sync"
  )
  
  func main(){
    par := make(chan int)
    impar := make(chan int)
    canal := make(chan bool) 
  
  go enviar(par,impar,canal)
   receber(par,impar,canal)
  

  }

  func enviar(p, i chan int,c chan bool) {
  //  x:= 100
    for a:=0; a <=199 ; a++{
      if a % 2 == 0{
        p <- a
      }else{
        i <- a
      }
    }
    close(p)
    close(i)
    c <- true
  }
  
  func receber(p, i chan int,c chan bool){
    for{
      select {
        case a := <- p:
          fmt.Println("O valor do canal par:", a)
        case a := <- i:
          fmt.Println("O valor do canal impar:", a)
        case d,ok := <- c:
          if !ok {
            fmt.Println("O valor:", d)
          }else{
        break
          }
      }
    } 
  }