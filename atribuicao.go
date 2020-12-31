package main

import (
  "fmt"
)

func main () {
  x := 28                       // := eu declaro e inicio uma variavel caso n defina uma a tipagem será automática
  y := "Minha idade é "
  fmt.Println("Hey!", y, x)            //print line
  fmt.Printf("X: %v, %T \n", x, x)        //print format
  fmt.Printf("Y: %v, %T\n\n\n", y, y)       //print format
    
  x = 38  //quando usa igual eu atribuo valor pq ele ja existe declarado
  fmt.Printf("X: %v, %T \n\n\n", x, x)        //print format

  x,z := 12,88  //eu consigo usar := pq o z ainda n foi declarado
  fmt.Printf("X: %v, %T \n", x, x) 
  fmt.Printf("X: %v, %T \n", z, z) 
  
  g := 10==10  // g foi declarada como boolean e durante o código g será booleano
  fmt.Println(g)
  
}

// := só é usado dentro do codeblocks
//pode pesquisar keyword para saber as palavras reservadas