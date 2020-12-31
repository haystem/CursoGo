package main

import ("fmt")

func main(){
  for i:=0; i<20;i++{
    if i%2 !=0{
    continue // continue divindo quando não for diferente 0 - ele para próxima interação - break quebra o loop todo -  continue quebra a interação daquele loop e passa para pròximo
    }
      fmt.Println(i)
  }
}
