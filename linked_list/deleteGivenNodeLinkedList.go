package linkedlist

// Delete given node in a Linked List : O(1) approach
// Problem Statement: Write a function to delete a node in a singly-linked list. You will not be given access to the head of the list instead, you will be given access to the node to be deleted directly. It is guaranteed that the node to be deleted is not a tail node in the list.
import "fmt"

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
