package linkedlist

import "fmt"

func LastNthNode() {
	n := 3
	link := List{}
	link.Insert(5)
	link.Insert(9)
	link.Insert(13)
	link.Insert(22)
	link.Insert(28)
	link.Insert(36)

	temp := link.head

	for i := 0; i < n-1; i++ {
		if temp.next == nil {
			fmt.Println("not possible")
			break
		}
		temp = temp.next
	}

	temp2 := link.head

	for temp.next != nil {
		temp = temp.next
		temp2 = temp2.next
	}

	fmt.Println("####################3")
	fmt.Println(temp2.key)

	link.Display()

}
