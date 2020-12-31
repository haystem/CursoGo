package main

import(
  "fmt"
  )
  
  type pessoa struct {
    nome string
    sobrenome string
    falar
  }
  
  type falar struct {
    falas []string
  }
  
  type humanos interface {
    fala()
  }
  
  func dizerAlgumaCoisa(x humanos) {
     x.fala()
  }
  
  
  func (p *pessoa) fala () {
        fmt.Println(p.nome);
        fmt.Println(p.sobrenome);
        fmt.Println(p.falar);
  }
  
  func main(){
    pessoa := pessoa{
       nome:"Alberto",
       sobrenome:"Santos",
       falar: falar{
         falas: []string{"Jesus", "Consegui depois de uma hora tentando","Esqueci que fiz como ponteiro"},
       },
    } 

    fmt.Println("\n *********************************************")
    dizerAlgumaCoisa(&pessoa)
  
  }
