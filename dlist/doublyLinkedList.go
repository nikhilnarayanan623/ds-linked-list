package dlist

import (
	"fmt"

	"github.com/nikhilnarayanan623/linkedlists/dlist/interfaces"
)

type node struct {
	data int
	next *node
	prev *node
}

type doublyLinkedList struct {
	head *node
	tail *node
}

// func to get new doubly linked list
func NewDoublyLinkedList() interfaces.DoublyLinkedListInterface {
	return &doublyLinkedList{}
}

// to add new vlaue to doubly linked list
func (d *doublyLinkedList) Addvalue(data int) {
	newNode := &node{data: data}

	if d.head == nil {
		d.head = newNode
	} else {
		d.tail.next = newNode // add a new value to the next of tail
		newNode.prev = d.tail //connext the tail to newNode (tail is now back of new node)
	}

	d.tail = newNode //set the last node as tail
}

// func to add multiple values to linked list
func (d *doublyLinkedList) AddMultipleValues() {
	var (
		limit int
		val   int
	)

	fmt.Print("how may values need to enter in Dlist: ")
	fmt.Scanf("%d", &limit)

	for i := 1; i <= limit; i++ {
		fmt.Print("Enter the value: ")
		fmt.Scanf("%d", &val)
		d.Addvalue(val)
	}
}

// insert a data after a specific
func (d *doublyLinkedList) InsertDataAfterAValue(checkVal int, data int) {

	//chekc the list is empty or not
	if d.head == nil {
		fmt.Println("doubly linked list is empty")
		return
	}

	newNode := &node{data: data} //create a new node

	//check the data is in head
	if d.head.data == checkVal {
		newNode.prev = d.head //connect new node to head
		//check head.next is nil or not if not then connect head next and new node
		if d.head.next != nil {
			newNode.next = d.head.next
			d.head.next.prev = newNode
		}
		d.head.next = newNode //connect head to new node
		return
	}

	//creat var to traverse through the list
	currentNode := d.head

	for currentNode != nil && currentNode.data != checkVal {
		currentNode = currentNode.next
	}
	//if value found, but not at tail
	if currentNode != nil && currentNode != d.tail {
		//connect new node and current nodes next
		newNode.next = currentNode.next
		newNode.next.prev = newNode

		//connect current node and new node
		currentNode.next = newNode
		newNode.prev = currentNode
		return
	}
	//value found at tail or value not found

	//connect new node and tail
	newNode.prev = d.tail
	d.tail.next = newNode

	//assign tail as newnode
	d.tail = newNode

	//if vlueu not found then show the message and connect new node to last of the list
	if currentNode == nil {
		fmt.Print("\n\ncant find the value in list so adding the data into last of the list\n\n")
	}
}

// func to remove a single value from doubly linked list
func (d *doublyLinkedList) RemoveSingleValue(value int) {

}

// func to print all values in forward order
func (d *doublyLinkedList) PrintAllForward() {
	//chekc the list is empty or not
	if d.head == nil {
		fmt.Println("doubly linked list is empty")
		return
	}
	// traverse head to tail
	currentNode := d.head

	fmt.Println("Doubly linked list values in forward order")
	for currentNode != nil {
		fmt.Print(currentNode.data, " ")
		//move on forwared
		currentNode = currentNode.next
	}
	fmt.Println()
}

// func to print all values in backward
func (d *doublyLinkedList) PrintAllBackward() {
	//chekc the list is empty or not
	if d.head == nil {
		fmt.Println("Doubly linked list is empty ")
		return
	}

	//travers the list tail to head
	currentNode := d.tail
	fmt.Println("Doubly linked list values in bakcaward order")

	for currentNode != nil {
		fmt.Print(currentNode.data, " ")
		currentNode = currentNode.prev
	}
	fmt.Println()
}
