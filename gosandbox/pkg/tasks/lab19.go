package tasks

import (
	"fmt"
)

type Node struct {
	name string
	time int
	dist float32
	next *Node
}

type Route struct {
	head   *Node
	length float32
}

func (r *Route) insertAtHead(name string, time int, dist float32) {
	nodeHead := &Node{name, time, dist, nil}
	if r.head == nil {
		r.head = nodeHead
	} else {
		nodeHead.next = r.head
		r.head = nodeHead
	}
	r.length += dist
}

func (r *Route) calcDistance(name1, name2 string) float32 {
	var dist float32
	var started bool
	node := r.head
	fmt.Println(node)
	for node != nil {
		fmt.Println("Traversing node", node.name)
		if started {
			dist += node.dist
			if node.name == name2 {
				break
			}
		}
		if node.name == name1 {
			started = true
		}
		node = node.next
	}
	return dist
}

func LabNineteen() {
	list := Route{nil, 0}
	list.insertAtHead("С", 15, 18)
	list.insertAtHead("B", 5, 6)
	list.insertAtHead("A", 0, 0)
	fmt.Println(list.calcDistance("A", "C"))
}
