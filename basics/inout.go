package main


import "fmt"

var name string
var age string

func main(){
	fmt.Print("Enter name: ")

	fmt.Scan(&name)
	fmt.Print("\n")
	fmt.Print("Enter age: ")
        fmt.Scan(&age)
	fmt.Print("\n")
        fmt.Println(name,age) 
}
