package main
import "fmt"
func main(){
	//array declartion
	var rollno[3]int
	//assign values
	rollno[0]=001
	rollno[1]=002
	rollno[2]=003
	fmt.Println("first roll_no",rollno[0])
	fmt.Println("last roll_no",rollno[2])
	fmt.Println("no of students",len(rollno))

	//slice
	names:=[]string{"angel","anu","ammu"}
	fmt.Println("names",names)
	//append
	names=append(names,"bincy")
	fmt.Println("new student added",names)
}