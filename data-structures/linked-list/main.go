package main

import (
	"fmt"
	"sync"
)

type Node struct {
	data int
	next *Node
}

type SinglyLinkedList struct {
	head *Node
	size int
	lock sync.RWMutex
}

// Adds an Item to the end of the list
func (sll *SinglyLinkedList) Append(item int) {
	sll.lock.Lock()
	node := Node{item, nil}

	if sll.head == nil {
		sll.head = &node
	} else {
		last := sll.head
		for {
			if last.next == nil {
				break
			}
			last = last.next
		}
		last.next = &node
	}
	sll.size++
	sll.lock.Unlock()

}

func (sll *SinglyLinkedList) Insert(pos int, item int) error {
	sll.lock.Lock()
	defer sll.lock.Unlock()
	if pos < 0 {
		return fmt.Errorf("Out of bounds")
	}
	newNode := Node{item, nil}
	// If insterting at head
	if pos == 0 || pos > sll.size {
		newNode.next = sll.head
		sll.head = &newNode
		return nil
	}

	// Every other pos
	currentNode := sll.head
	j := 0
	for j < pos-2 {
		j++
		currentNode = currentNode.next
	}
	newNode.next = currentNode.next
	currentNode.next = &newNode
	sll.size++
	return nil
}

func (sll *SinglyLinkedList) Print() {
	sll.lock.RLock()
	defer sll.lock.RUnlock()
	node := sll.head
	for {
		if node == nil {
			break
		}
		fmt.Print(node.data)
		node = node.next
		if node == nil {
			break
		}
		fmt.Print("->")

	}
	fmt.Println()
}

func main() {
	list := SinglyLinkedList{}
	err := list.Insert(3, 33)
	if err != nil {
		fmt.Print(err)
	}
	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.Append(5)
	list.Append(44)
	list.Print()
	fmt.Printf("Size: %d , Type: %T ", list.size, list.size)
}
