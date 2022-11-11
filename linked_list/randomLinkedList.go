package linkedlist

import "fmt"

type RandomNode struct {
	next   *RandomNode
	random *RandomNode
	key    interface{}
}

func RandomLinkedList() {
	input := [][]interface{}{{7, nil}, {13, 0}, {11, 4}, {10, 2}, {1, 0}}
	arr := []RandomNode{}
	for i := 0; i < len(input); i++ {
		node := RandomNode{}
		node.key = input[i][0]
		arr = append(arr, node)
	}

	for i := 0; i < len(arr); i++ {
		if i < len(arr)-1 {
			arr[i].next = &arr[i+1]
		}
		j, ok := input[i][1].(int)
		if ok {
			arr[i].random = &arr[j]
		}
	}
	head := &arr[0]
	arr = nil
	input = nil
	temp := head
	for {
		// fmt.Println(temp.key)
		if temp.next == nil {
			break
		}
		temp = temp.next
	}
	// fmt.Printf("%#v\n", arr)
	// fmt.Printf("%#v\n", input)

	// fmt.Println(len(arr), cap(arr))
	// fmt.Println(len(input), cap(input))

	temp = head
	for {
		node := RandomNode{}
		if temp.next == nil {
			node.next = nil
			node.key = temp.key
			temp.next = &node
			break
		}
		node.next = temp.next
		node.key = temp.key
		// fmt.Println("before", temp, node)
		temp.next = &node
		if node.next == nil {
			break
		}
		// fmt.Println("after", temp, node)
		temp = node.next
	}
	orignal := head
	result := head.next
	output := head.next

	for {
		if result.next == nil {
			break
		}
		orignal.next = orignal.next.next
		orignal = orignal.next
		result.next = orignal.next
		result = result.next
	}

	temp = output
	for {
		fmt.Println(temp)
		if temp.next == nil {
			break
		}
		temp = temp.next
	}

	// for {
	// 	fmt.Println(head.key)
	// 	if head.random == nil {
	// 		break
	// 	}
	// 	head = *head.random
	// }

}
