package stackandqueue

import "fmt"

type Stack struct {
	element []int
	len     int
}

func (s *Stack) push(e int) {
	if s.isFull() {
		fmt.Println("Overflow")
	} else {
		s.element[s.len] = e
		s.len++
	}
}

func (s *Stack) pop() int {

	if s.isEmpty() {
		fmt.Println("Underflow")
		return -1
	}
	element := s.element[s.len-1]
	s.len--
	return element
}

func (s *Stack) peek() int {
	if s.isEmpty() {
		fmt.Println("Stack is Empty")
		return -1
	}
	return s.element[s.len-1]
}

func (s *Stack) isEmpty() bool {
	if s.len <= 0 {
		return true
	}
	return false
}

func (s *Stack) isFull() bool {
	if s.len >= 15 {
		return true
	}
	return false
}

type Queue struct {
	input  Stack
	output Stack
}

func (q *Queue) enqueue(e int) {
	q.input.push(e)
}

func (q *Queue) dequeue() int {
	if q.output.isEmpty() {
		for {
			if q.input.len <= 1 {
				break
			}
			q.output.push(q.input.pop())
		}
		return q.input.pop()
	}
	return q.output.pop()
}

func QueueUsingStack() {

	input := Stack{
		element: []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		len:     0,
	}
	output := Stack{
		element: []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		len:     0,
	}

	q := Queue{
		input:  input,
		output: output,
	}
	for i := 1; i < 8; i++ {
		q.enqueue(i)
	}
	for i := 0; i < 2; i++ {
		fmt.Println(q.dequeue())
	}
	for i := 100; i < 103; i++ {
		q.enqueue(i)
	}
	for i := 0; i < 10; i++ {
		fmt.Println(q.dequeue())
	}
}
