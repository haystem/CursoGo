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
     x.fala();
     x.nomeFala();
     x.sobrenomeFala();
  }
  
  
  func (p pessoa) fala() {
        fmt.Println(p.nome);
        fmt.Println(p.sobrenome);
        fmt.Println(p.falar);
  }
  
    func (p pessoa) nomeFala() {
        fmt.Println(p.nome);
        fmt.Println(p.falar);
  }
  
    func (p pessoa) sobrenomeFala() {
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
  
      pessoa1 := pessoa{
       nome:"Alberto",
       falar: falar{
         falas: []string{"Jesus", "Consegui depois de uma hora tentando","Esqueci que fiz como ponteiro"},
       },
      }
    
      pessoa2 := pessoa{
       sobrenome:"Santos",
       falar: falar{
       falas: []string{"Jesus", "Consegui depois de uma hora tentando","Esqueci que fiz como ponteiro"},
       },
      }
    
  
    dizerAlgumaCoisa(pessoa)
    
    fmt.Println("\n *********************************************")
    
        dizerAlgumaCoisa(pessoa1)
    
    fmt.Println("\n *********************************************")
    
        dizerAlgumaCoisa(pessoa2)
    
    fmt.Println("\n *********************************************")
  
  }
