package main

import(
  "fmt"
  "log"
)

type Erro struct{
  anyError string
}

func (e Erro) Error() string{
   return fmt.Sprintf("O erro encontrado foi %v : ", e.anyError)
}

func main (){

 s, err := sqrt(0)
 if err != nil{
   log.Fatalln(err)
 }
 defer fmt.Println(s)
}

func sqrt(f float64) (float64, error){
  if f <= 0{
    return 0,Erro{"ola"}
  }
  return 42,nil
}