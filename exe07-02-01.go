package main 

import ("fmt")

type pessoa struct {
  nome string
  sobrenome string
  idade int
}

func main(){
  pess := pessoa{"alberto","rocha",30}
  fmt.Println(pess)
  mudeMe(&pess)
  fmt.Println(pess)
}

func mudeMe(p *pessoa) {
  (*p).nome = "Jussara"
  (*p).sobrenome = "Jussara"
  (*p).idade = 19
}