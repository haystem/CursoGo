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
  
  func (p *pessoa) fala (nm string, sb string, flr []string ) {
        fmt.Println(p.nome, nm);
        fmt.Println(p.sobrenome, sb);
        fmt.Println(p.falar,flr);
  }
  
  func main(){
    pessoa := pessoa{
       nome:"Alberto",
       sobrenome:"Santos",
       falar: falar{
         falas: []string{"Jesus", "O"},
       },
    } 
    
    pessoa.fala("Claudio","Jose",[]string{"Doria","Madureira"})
    
  }