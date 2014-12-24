package main

import (
"fmt"
)

func main(){
	withThisToken:=`xoxp-3057259082-3057259092-3247338018-9890ee`
  if socket,err:=getMeASocket(withThisToken); err != nil{
      fmt.Println(err)
	}
}
