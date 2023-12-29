package main

import (
	"fmt"
	"os"
	"strings"
)

//func Join(elems []string, sep string) string
func main(){
 	fmt.Println(strings.Join(os.Args[1:], " "))
}	
