package main
import "fmt"
func add(a int,b int)int{
	return a+b
}
func sub(a int,b int)int{
	return a-b
}
func main(){
	add:=add(10,20)
	fmt.Println("addition result:",add)
	substract:=sub(10,20)
	fmt.Println("substraction result:",substract)
}