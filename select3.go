package main

import("fmt")

//- Chans par, ímpar, quit 
//- Func send manda números pares pra um canal, ímpares pra outro, e fecha/quit 
//- Func receive é um select entre os três canais, encerra no quit
//- Problema!

func main(){
   par := make(chan int)
   impar := make(chan int)
   quit := make(chan bool)
   
  go enviar(par,impar,quit)
   receber(par,impar,quit)
  
}

func enviar(par, impar chan int, quit chan bool){
  for i:=0; i < 300; i++{
    if i%2 ==0{
      par <- i
    }else{
      impar <- i
    }
  }
  close(par)
  close(impar)
    quit <- true
}

func receber(par, impar chan int, quit chan bool){
  for{
  select {
    case c := <- par:
      fmt.Println("O valor: ",c,"é par!")
    case c := <- impar:
      fmt.Println("O valor: ",c,"é impar!")
    case c,ok := <- quit:
    if !ok{
        fmt.Println("Deu zebra! ", c)
     }else{
    fmt.Println("Encerrando :", c)
     }
    return
   }
 }
}