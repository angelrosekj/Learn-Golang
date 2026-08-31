package main
import "fmt"
func main(){
	age:=23
	p:=&age
	fmt.Println("value of age:",age)
	fmt.Println("address of age:",&age)
	fmt.Println("Pointer p:",p)
	fmt.Println("value using pointer:",*p)
}