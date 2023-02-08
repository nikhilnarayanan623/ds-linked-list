package dlist

import (
	"fmt"

	"github.com/nikhilnarayanan623/ds-linked-list/dlist/interfaces"
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
func (d *doublyLinkedList) Append(data int) {
	newNode := &node{data: data}

	if d.head == nil {
		d.head = newNode
	} else {
		d.tail.next = newNode // add a new value to the next of tail
		newNode.prev = d.tail //connext the tail to newNode (tail is now back of new node)
	}

	d.tail = newNode //set the last node as tail
}

// insert a vlaue to the beging of linked list
func (d *doublyLinkedList) Prepend(data int) {

	newNode := &node{data: data}

	newNode.next = d.head

	//check head is nil or not
	if d.head != nil {
		d.head.prev = newNode
	}
	//set newNode as head
	d.head = newNode

	//if list have only one value after append then assign tail as also head
	if d.head.next == nil {
		d.tail = d.head
	}
}

// func to add multiple values to linked list
func (d *doublyLinkedList) AppendMultipleValues() {
	var (
		limit int
		val   int
	)

	fmt.Print("how may values need to enter in Dlist: ")
	fmt.Scanf("%d", &limit)

	for i := 1; i <= limit; i++ {
		fmt.Print("Enter the value: ")
		fmt.Scanf("%d", &val)
		d.Append(val)
	}
}

// insert a data after a specific
func (d *doublyLinkedList) InsertAfterAValue(checkVal int, data int) {

	//chekc the list is empty or not
	if d.head == nil {
		fmt.Println("doubly linked list is empty so can't insert a value after specified")
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

	//check value found
	if currentNode == nil {
		fmt.Println("Value not found in doubly linked list so returning nothing will change")
		return
	}

	//connect newNode next and prev
	newNode.next = currentNode.next
	newNode.prev = currentNode
	//connect current next as newNode
	currentNode.next = newNode //if its nil or not not a problem
	//if current node is tail then set newNode as tail and return
	if currentNode == d.tail {
		d.tail = newNode
		return
	}
	//other wise connect newNode next's prev to newNode that we got from currentNode
	newNode.next.prev = newNode
}

func (d *doublyLinkedList) InsertBeforeAValue(checkVal int, data int) {
	//chekc the list is empty or not
	if d.head == nil {
		fmt.Println("doubly linked list is empty so can't insert a value after specified")
		return
	}

	newNode := &node{data: data}

	//check the value in head
	if d.head.data == checkVal {
		newNode.next = d.head
		d.head.prev = newNode
		d.head = newNode
		return
	}
	// loop to find the value
	currentNode := d.head.next
	for currentNode != nil && currentNode.data != checkVal {
		currentNode = currentNode.next
	}

	//check value found
	if currentNode == nil {
		fmt.Println("Value not found in doubly linked list so returning nothing will change")
		return
	}

	// in here we don't need to worry about tail becuase connection won't affect tail

	//connect currentNode prev's node and newNode
	newNode.prev = currentNode.prev
	currentNode.prev.next = newNode

	//connect currentNode and newNode
	newNode.next = currentNode
	currentNode.prev = newNode
}

// func to remove a single value from doubly linked list
func (d *doublyLinkedList) DeleteValue(value int) {

	if d.head == nil {
		fmt.Println("The doubly linked list is empty can't remove data")
		return
	}

	//check the value need to remove is in head or not
	if d.head.data == value {

		//check head next have value is available or not
		if d.head.next != nil {
			d.head.next.prev = nil //disconnect head next's prev connection
		}
		d.head = d.head.next //assign head as its next
		return
	}

	//value not in head then traverse through list to find the value
	currentNode := d.head.next

	for currentNode != nil && currentNode.data != value {
		currentNode = currentNode.next
	}

	if currentNode == nil {
		fmt.Println("Entered value to remove from the doubly linked list not found")
		return
	}

	// first consider the value not found at tail
	if currentNode != d.tail {

		currentNode.prev.next = currentNode.next
		currentNode.next.prev = currentNode.prev
		return
	}

	//if value found at tail then tail as tail prev and disconnect next

	d.tail = d.tail.prev
	d.tail.next = nil
}

// func to print all values in forward order
func (d *doublyLinkedList) DisplayOrder() {
	//chekc the list is empty or not
	if d.head == nil {
		fmt.Println("doubly linked list is empty")
		return
	}
	// traverse head to tail
	currentNode := d.head

	fmt.Println("Doubly linked list values in  Order")
	for currentNode != nil {
		fmt.Print(currentNode.data, " ")
		//move on forwared
		currentNode = currentNode.next
	}
	fmt.Println()
}

// func to print all values in backward
func (d *doublyLinkedList) DisplayReverse() {
	//chekc the list is empty or not
	if d.head == nil {
		fmt.Println("Doubly linked list is empty ")
		return
	}

	//travers the list tail to head
	currentNode := d.tail
	fmt.Println("Doubly linked list values in Reverse order")

	for currentNode != nil {
		fmt.Print(currentNode.data, " ")
		currentNode = currentNode.prev
	}
	fmt.Println()
}

//func to inset an array values to linked list

func (d *doublyLinkedList) ArrayToSlist(array []int) {

	for _, val := range array {
		d.Append(val)
	}
}

// func to remove duplicates if list in a sorted way
func (d *doublyLinkedList) RemoveDuplicates() {

	if d.head == nil {
		fmt.Println("Doubly linked list is empty ")
		return
	}

	prev := d.head
	currentNode := d.head.next

	for currentNode.next != nil {

		if currentNode.data == prev.data { //same value then move
			currentNode = currentNode.next
			continue
		}
		//connect the different value nodes
		prev.next = currentNode
		currentNode.prev = prev

		//set prev as currentNode
		prev = currentNode
	}

	//consider this 1 2 3 3 3
	// currentNode will be nil and prev will be 3
	// need prev as tail and disconnect all other 3s

	if currentNode.data == prev.data {
		prev.next = nil
		d.tail = prev
	}
}
