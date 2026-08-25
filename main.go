package main
import "fmt"
func main(){
	fmt.Println("hello world")
	
	//variables(explicitly specify data type)
	var name string="angel"//string data type
	var age int=23//integer
	 
	//go infer the data type
	mark:=20
	passed:=false//bool
	letter:='a'//byte

	//printing variables declared ,otherwise error
	fmt.Println("Name:",name)
	fmt.Println("Age:",age)
	fmt.Println("Mark:",mark)
	fmt.Println("Passed:",passed)
	fmt.Println("Letter:",letter)

	//constants for values that donot change
	const dob=2003
	fmt.Println("DOB:",dob)

	//diffrence between printf, println and print
	fmt.Print("dare to dream")
	fmt.Print("angel")
    //new line
	fmt.Println("dare to dream")
	fmt.Println("angel")
	//format output
	fmt.Printf("my name is %s and i am %d years old",name,age)
    fmt.Println("")
	//if else 
	if age>=18{
		fmt.Println("angel is eligible for vote")
	}else{
		fmt.Println("angel is not eligibile for vote")
	}
	switch{
	case mark >80:
		fmt.Println("grade:A+")
		fallthrough
	case mark >70:
		fmt.Println("grade:B+")
		fallthrough//tells to execute next case
	default:
		fmt.Println("failed")
	}
}

	

	


