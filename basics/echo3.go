//simple go echo impl. from book
package main



import (
	"fmt"
	"os"
)

func main(){
	fmt.Print(os.Args[0] + ": ")
	var s, sep string
	for i,arg := range os.Args[1:] {
		fmt.Println(i,arg)
		s += sep + arg;
		sep = " "
}

	fmt.Println(s)
}
