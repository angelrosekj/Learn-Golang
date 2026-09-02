package main

import "fmt"

type Queue struct {
	Items [5]string
	Size  int
}

func (q *Queue) Enqueue(element string) {
	if q.IsFull() {
		fmt.Println("Queue is full")
		return
	}

	q.Items[q.Size] = element
	q.Size++

	fmt.Println("Element added:", element)
}

func (q *Queue) Dequeue() {
	if q.IsEmpty() {
		fmt.Println("Queue is empty")
		return
	}

	fmt.Println("Element removed:", q.Items[0])

	// Move elements one position forward
	for i := 0; i < q.Size-1; i++ {
		q.Items[i] = q.Items[i+1]
	}

	q.Items[q.Size-1] = ""
	q.Size--
}

func (q Queue) IsEmpty() bool {
	return q.Size == 0
}

func (q Queue) IsFull() bool {
	return q.Size == len(q.Items)
}

func main() {
	queue := Queue{}

	queue.Enqueue("Nikhitha")
	queue.Enqueue("Anu")
	queue.Enqueue("Riya")

	fmt.Println(queue.Items)

	queue.Dequeue()

	fmt.Println(queue.Items)

	queue.Dequeue()

	fmt.Println(queue.Items)
}