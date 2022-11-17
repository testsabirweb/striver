package stackandqueue

import "fmt"

type Node struct {
	next     *Node
	previous *Node
	key      int
}

type List struct {
	head   *Node
	end    *Node
	length int
}

type LRU struct {
	link  List
	dict  map[int]*Node
	limit int
}

func (l *List) addNode(k int) {
	//list is empty
	if l.head == nil {
		l.head = &Node{
			next:     nil,
			previous: nil,
			key:      k,
		}
		l.end = l.head
		l.length = 1
	} else {
		node := &Node{
			next:     nil,
			key:      k,
			previous: nil,
		}
		l.head.previous = node
		node.next = l.head
		l.head = node
		l.length = l.length + 1
	}
}

func (l *List) removeNode(n *Node) {
	if n.next == nil { //if node is at the last
		l.end = n.previous
		if l.end != nil {
			l.end.next = nil
		} else {
			l.head = nil
		}

	} else if n.previous == nil { // if node is begning
		l.head = n.next
		l.head.previous = nil
	} else {
		next, previous := n.next, n.previous
		previous.next = next
		next.previous = previous
	}
	n = nil
	l.length = l.length - 1
}

func (l *List) Display() {
	list := l.head
	for list != nil {
		fmt.Printf("%+v ->", list.key)
		list = list.next
	}
	fmt.Println()
}

func (lru *LRU) insertLRU(k int) {
	if lru.link.length >= lru.limit {
		lru.dict[lru.link.end.key] = nil
		lru.link.removeNode(lru.link.end)
	}
	lru.link.addNode(k)
	lru.dict[k] = lru.link.head
	lru.link.Display()
}

func (lru *LRU) getLRU(k int) int {
	if lru.dict[k] == nil {
		fmt.Println(-1, k, "########")
		return -1
	}
	node := lru.dict[k]
	lru.link.removeNode(node)
	lru.link.addNode(k)
	lru.dict[k] = lru.link.head
	lru.link.Display()
	return k
}

func LRUCache() {
	lru := LRU{
		limit: 2,
		dict:  map[int]*Node{},
		link:  List{},
	}
	commands := []string{"put", "put", "get", "put", "get", "put", "get", "get", "get"}
	inputs := [][]int{{1, 1}, {2, 2}, {1}, {3, 3}, {2}, {4, 4}, {1}, {3}, {4}}
	for index, c := range commands {
		if c == "put" {
			lru.insertLRU(inputs[index][0])
		} else {
			lru.getLRU(inputs[index][0])
		}
	}

	// dict := map[int]Node{}
}
