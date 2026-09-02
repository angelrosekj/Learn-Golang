package main

import "fmt"

type StackArray struct {
	items [5]string
	size  int
}

func (arr *StackArray) Push(element string) {
	if arr.IsFull() {
		fmt.Println("Stack is full")
		return
	}

	arr.items[arr.size] = element
	arr.size++

	fmt.Println("String is added to stack:", arr.items)
}

func (arr *StackArray) Pop() {
	if arr.IsEmpty() {
		fmt.Println("Stack is empty, cannot pop")
		return
	}

	arr.size--
	fmt.Println("Last element is deleted:", arr.items[arr.size])

	arr.items[arr.size] = ""
}

func (arr *StackArray) IsEmpty() bool {
	if arr.size == 0 {
		return true
	}
	return false
}

func (arr *StackArray) IsFull() bool {
	if arr.size == len(arr.items) {
		return true
	}
	return false
}

func (arr *StackArray) Peek() {
	if arr.IsEmpty() {
		fmt.Println("Stack is empty")
		return
	}

	fmt.Println("Top element:", arr.items[arr.size-1])
}

func main() {
	stack1 := StackArray{}

	stack1.Push("jack")
	stack1.Push("alan")

	stack1.Pop()

	stack1.Peek()
}