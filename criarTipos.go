package main 

import (
  "fmt"
  )

type hotdog int // int é o tipo subjacente 

var b hotdog = 10
  
  func main(){
    a:= 10  // o tipo de b não é igual ao valor de a pois, B= hotdog e A= int
     fmt.Printf(" o valor de a = %v , e o tipo %T \n\n ",a,a)
    fmt.Printf(" o valor de b = %v , e o tipo %T \n\n", b,b)
    
    z:= int(b)
      fmt.Printf(" o valor de z convertendo b = %v , e o tipo %T \n\n ", z,z)
    b = 30303030
    
      fmt.Printf(" o valor de b = %v , e o tipo %T \n\n", b,b)
      
     /*   b = a // não é possivel pelo tipo divergente
      fmt.Printf(" o valor de b = %v , e o tipo %T \n\n", b,b) */
  }
