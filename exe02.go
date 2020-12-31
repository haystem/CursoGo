package main;

import (
  "fmt"
  );
  
  var x int ;
  var y string ;
  var z bool;
  
  func main (){
     x = 42;
     y = "James Bond";
     z = true;
     s := fmt.Sprintf("O valor de X, Y e Z nas respectiva ordem são: %v, %v e %v",x,y,z) ;
     
     fmt.Println("Na váriavel S contem: ",s);
  }