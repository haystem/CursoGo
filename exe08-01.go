package main

import (
	"fmt"
	"encoding/json"
	"os"
)

type User struct {
	First string
	Age   int
}

func main() {
	u1 := User{
		First: "James",
		Age:   32,
	}

	u2 := User{
		First: "Moneypenny",
		Age:   27,
	}

	u3 := User{
		First: "M",
		Age:   54,
	}

	users := []User{u1, u2, u3}
	
	jsonFinal,err:= json.Marshal(users)
  if err != nil{fmt.Println(err)}
  fmt.Println(string(jsonFinal))
  os.Stdout.Write(jsonFinal)

	// your code goes here
}
