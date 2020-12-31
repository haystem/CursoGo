package main 

import (
  "fmt"
  )
  var x int = 10  //usar apenas quando tiver fora do codeblocks
  var y string = "Isso é uma string" // declarado fora
  var z int // o valor só pode ser atribuido dentro de um codeblocks 
  
  // variavel global não precisa ser usado
  
  func main(){
    // o tipo declarado é para sempre no programa - tipagem estática
    x = 20
    fmt.Println(x)
    
   //  x = 20.5 // não dá pq mudaria o tipo de int para float e a linguagem é fortemente tipada
  }
