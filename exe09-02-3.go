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
    nomeFalas()
    sobrenomeFala()
  }
  
  func dizerAlgumaCoisa(x humanos) {
     x.fala()
  }
  
  
  func (p pessoa) fala () {
        fmt.Println(p.nome);
        fmt.Println(p.sobrenome);
        fmt.Println(p.falar);
  }
    func (p pessoa) nomeFalas () {
        fmt.Println(p.nome);
        fmt.Println(p.falar);
  }
  
    func (p pessoa) sobrenomeFala () {
        fmt.Println(p.sobrenome);
        fmt.Println(p.falar);
  }
  
  func main(){
    pessoa := pessoa{
       nome:"Alberto",
       sobrenome:"Santos",
       falar: falar{
         falas: []string{"Jesus", "O"},
       },
    } 
    
    
    pessoa.fala()
    
    fmt.Println("\n *********************************************")
    
  dizerAlgumaCoisa(pessoa)
    
  }