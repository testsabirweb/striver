package linkedlist

// Delete given node in a Linked List : O(1) approach
// Problem Statement: Write a function to delete a node in a singly-linked list. You will not be given access to the head of the list instead, you will be given access to the node to be deleted directly. It is guaranteed that the node to be deleted is not a tail node in the list.
import "fmt"

type Node struct {
	prev *Node
	next *Node
	key  interface{}
}

type List struct {
	head *Node
	tail *Node
}

func (L *List) Insert(key interface{}) {
	list := &Node{
		next: L.head,
		key:  key,
	}
	if L.head != nil {
		L.head.prev = list
	}
	L.head = list

	l := L.head
	for l.next != nil {
		l = l.next
	}
	L.tail = l
}

func (l *List) Display() {
	list := l.head
	for list != nil {
		fmt.Printf("%+v ->", list.key)
		list = list.next
	}
	fmt.Println()
}

func Display(list *Node) {
	for list != nil {
		fmt.Printf("%v ->", list.key)
		list = list.next
	}
	fmt.Println()
}

func ShowBackwards(list *Node) {
	for list != nil {
		fmt.Printf("%v <-", list.key)
		list = list.prev
	}
	fmt.Println()
}

func (l *List) Reverse() {
	curr := l.head
	var prev *Node
	l.tail = l.head

	for curr != nil {
		next := curr.next
		curr.next = prev
		prev = curr
		curr = next
	}
	l.head = prev
	Display(l.head)
}

func DeleteGivenNodeLinkedList() {
	link := List{}
	link.Insert(5)
	link.Insert(9)
	link.Insert(13)
	link.Insert(22)
	link.Insert(28)
	link.Insert(36)

	fmt.Println("\n==============================\n")
	fmt.Printf("Head: %v\n", link.head.key)
	link.Display()

	givenNode := link.head.next.next
	fmt.Println(givenNode.key, "#########")
	for {
		if givenNode.next.next == nil {
			break
		}
		givenNode.key = givenNode.next.key
		givenNode = givenNode.next
	}
	givenNode.key = givenNode.next.key
	givenNode.next = nil

	fmt.Println("\n==============================\n")
	fmt.Printf("Head: %v\n", link.head.key)
	link.Display()
}
