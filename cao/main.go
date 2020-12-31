//Cachorro é um pacote que retorna em anos humanos para cães
package cao

import (
  "fmt"
 // "log"
  )

var (i int 
    c int
    )
//Idade pergunta o número e joga para variavel
func Idade() int{
   fmt.Println("Digite a Idade:")
   fmt.Scan(&i)
   return i
}
//Converte o valor multiplicado para 7  
func Cao(d int){
  c:= d*7
  fmt.Println("A idade é :", c)
}
//Chame função com pergunta do ano escolhendo qualquer variedade
func Conv (){
   Idade()
  defer Cao(i)
}