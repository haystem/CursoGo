package main 
import (
  "fmt" 
  "github.com/yosssi/gohtml"
  ) 
  func main(){ 
    
    h := `<!DOCTYPE html><html><head><title>This is a title.</title><script type="text/javascript"> alert('aaa'); if (0 < 1) { 	alert('bbb'); } </script><style type="text/css"> body {font-size: 14px;} h1 { 	font-size: 16px; 	font-weight: bold; } </style></head><body><form><input type="name"><p>AAA<br>BBB></p></form><!-- This is a comment. --></body></html>` 	
    
    fmt.Println(gohtml.Format(h))
    
    
  }