package linkedlist

import "fmt"

type node struct {
	data int
	next *node
}

func SinglyLinkedList() {

	head := &node{
		data: 2,
		next: nil,
	}
	head.next = &node{
		data: 3,
		next: nil,
	}
	head.next.next = &node{
		data: 4,
		next: nil,
	}
	head.next.next.next = &node{
		data: 5,
		next: nil,
	}

	temp := head

	for {
		if temp.next == nil {
			break
		}
		fmt.Println(temp.data)
		temp = temp.next
	}
}
