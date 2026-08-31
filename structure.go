package main
import "fmt"
type student struct{
	name string
	age int
	mark float32
}
func main(){
	student1:=student{}
	student1.name="angel"
	student1.age=23
	student1.mark=34.2
	fmt.Println("name of student",student1.name)
    fmt.Println("age of student",student1.age)
	fmt.Println("mark of student",student1.mark)
}