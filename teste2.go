package main 

import(
  "fmt")
  
func main() {
  var rep, rep2,rep3 string
  fmt.Print("Nome: ")
  fmt.Scan(&rep3)
  fmt.Print("Sobrenome: ")
    fmt.Scan(&rep2)
  fmt.Print("Apelido: ")
    fmt.Scan(&rep)
  
  fmt.Println(rep3,rep2,rep)
  
}