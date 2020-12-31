package main

import(
  "fmt"
  "golang.org/x/crypto/bcrypt"
  )

func main(){
 
 passwd:="Perebas"
 
 hash,err := bcrypt.GenerateFromPassword([]byte(passwd), 10)
  if err != nil{
    fmt.Println(err)
  }
  fmt.Println(hash)
  fmt.Println(string(hash),"Convertido para String")
}