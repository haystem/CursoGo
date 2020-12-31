/* OS zeros

int = 0
float = 0.0 
booleans = false
strings = ""
pointers, functions, interfaces, slices, channels, maps = nil */

package main
import (
  "fmt"
  )
  
  var a int
  var b bool
  var c float32
  var d string
  
  func main(){
    fmt.Printf("%v, %T\n",a,a)
    fmt.Printf("%v, %T\n",b,b)
    fmt.Printf("%v, %T\n",c,c)
    fmt.Printf("%v, %T\n",d,d)
  }