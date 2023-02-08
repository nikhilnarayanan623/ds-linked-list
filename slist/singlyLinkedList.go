package slist

import (
	"fmt"

	"github.com/nikhilnarayanan623/ds-linked-list/slist/interfaces"
)

// for a single value in linked list
type node struct {
	data int
	next *node
}

// to hold head and tail of linked list and its functions
type singlyLinkedlist struct {
	head *node
	tail *node
}

// to get a new linked list
func NewSinglyLinkedList() interfaces.LinkedListInterface {

	return &singlyLinkedlist{}
}

// to add new value to linked list
func (l *singlyLinkedlist) Append(data int) {
	//create a new node
	newNode := &node{data: data}
	if l.head == nil { //if head is nil then assign new node to head
		l.head = newNode
	} else {
		l.tail.next = newNode
	}
	l.tail = newNode //assign tail as new node for next time add values
}

//add a value to the starting of linked list

func (l *singlyLinkedlist) Prepend(data int) {

	newNode := &node{data: data}

	//connect new node and head
	newNode.next = l.head
	l.head = newNode

	//if list have no extr value then make it as tail also
	if l.head.next == nil {
		l.tail = l.head
	}
}

// add multiple values
func (l *singlyLinkedlist) AppendMultipleValues() {
	var (
		limit int
		val   int
	)

	fmt.Print("how may values need to enter in Slist: ")
	fmt.Scanf("%d", &limit)

	//loop how many values want
	for i := 1; i <= limit; i++ {
		fmt.Print("Enter value: ")
		fmt.Scanf("%d", &val) //get the value from user

		l.Append(val) //pass the value to insert int linked list
	}
}

// func insert before a value
func (l *singlyLinkedlist) InsertBeforeAValue(checkVal int, data int) {

	if l.head == nil {
		fmt.Println("Singly linked list is empty")
		return
	}

	newNode := &node{data: data}
	//check the values is contain on head

	if l.head.data == checkVal {
		newNode.next = l.head
		l.head = newNode
		return
	}

	//find the node as next node
	cuurentNode := l.head

	for cuurentNode.next != nil && cuurentNode.next.data != checkVal {
		cuurentNode = cuurentNode.next
	}

	if cuurentNode.next == nil {
		fmt.Println("can't find the value in linked list")
		return
	}
	//value found at currentNode's next node
	newNode.next = cuurentNode.next
	cuurentNode.next = newNode
}

// func to insert a value after a specific values in the list
func (l *singlyLinkedlist) InsertAfterAValue(checkVal int, data int) {
	//chekc the list is empty or not
	if l.head == nil {
		fmt.Println("Singly linked list is empty")
		return
	}

	//create a new node
	newNode := &node{data: data}

	//check the values is in head
	if l.head.data == checkVal {
		newNode.next = l.head.next
		l.head.next = newNode
		return
	}

	currentNode := l.head

	for currentNode != nil && currentNode.data != checkVal {
		currentNode = currentNode.next
	}

	//check value found or not
	if currentNode == nil {
		fmt.Println("Value that given to check is not founded so nothing will change")
		return
	}

	newNode.next = currentNode.next
	currentNode.next = newNode
	//if currentNode is tail then change tail as newNode
	if currentNode == l.tail {
		l.tail = currentNode
	}
}

// to remove values from linked list
func (l *singlyLinkedlist) DeleteValue(value int) {
	currentNode := l.head
	//check if the linked list is empty or not
	if currentNode == nil {
		fmt.Println("There is no values in the list")
		return
	}

	//check the value is in head
	if currentNode.data == value {
		l.head = currentNode.next
		return
	}

	//loop and find the value
	for currentNode.next != nil {
		if currentNode.next.data == value {
			//chekc the currentNode.next is in tail or not
			if currentNode.next == l.tail {
				l.tail = currentNode //current node as tail
			}
			currentNode.next = currentNode.next.next //assign current node next as its next's next so we can remove the next value
			return
		}
		currentNode = currentNode.next
	}
	fmt.Printf("%d is not present in linked list", value)
}

// print all values forward
func (l *singlyLinkedlist) DisplayOrder() {
	newNode := l.head

	if newNode == nil {
		fmt.Println("Linked list is empty")
		return
	}
	fmt.Println("\nLinked list values in Order")
	for newNode != nil {
		fmt.Print(newNode.data, " ")
		newNode = newNode.next
	}
	fmt.Println()
}

// to print all values in backward
func (l *singlyLinkedlist) DisplayReverse() {

	currentNode := l.head

	if currentNode == nil {
		fmt.Println("Linked list is empty")
		return
	}
	fmt.Println("\nLinked list values in Reverse order")

	DisplayReverseHelper(currentNode)
	fmt.Println()
}

// func to help displayOrder in recurrsion
func DisplayReverseHelper(node *node) {

	if node == nil {
		return
	}

	DisplayReverseHelper(node.next)
	fmt.Print(node.data, " ")
}

//functions to remove duplicate values from singly linked list

func (l *singlyLinkedlist) RemoveDuplicates() {

	if l.head == nil {
		fmt.Println("The signly linked list is empty")
		return
	}

	prevNode := l.head
	currentNode := l.head.next

	for currentNode != nil {

		if currentNode.data == prevNode.data {
			currentNode = currentNode.next
			continue
		}

		prevNode.next = currentNode

		prevNode = currentNode
		currentNode = currentNode.next
	}

}

// array to linked list

func (l *singlyLinkedlist) ArrayToSlist(array []int) {

	for _, val := range array {
		l.Append(val)
	}
}
