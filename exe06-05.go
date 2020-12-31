package main

import (
    "fmt"
  )
  
  // incluir todos os structs
  
  type quadrado struct {
    lado float64
  }
  type circulo struct {
    raio float64
  }
  
  
  func main(){
    
    s1:= quadrado{lado :6}
    s2:=circulo{raio:87} 
    
    info(s1)
    info(s2)
    }
  
  // metodos para area
  
  func (l quadrado) area() float64{
    return l.lado*l.lado //area quadrado
  }
  
  func (r circulo) area() float64{
    return 2*(r.raio * r.raio)
  }
  
  type figura interface {
    area() float64
  }
  
  func info(f figura) {
    fmt.Println(f.area())
  }