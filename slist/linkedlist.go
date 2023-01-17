package slist

import (
	"fmt"

	"github.com/nikhilnarayanan623/linkedlists/slist/interfaces"
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
func (l *singlyLinkedlist) Addvalue(data int) {
	//create a new node
	newNode := &node{data: data}
	if l.head == nil { //if head is nil then assign new node to head
		l.head = newNode
	} else {
		l.tail.next = newNode
	}
	l.tail = newNode //assign tail as new node for next time add values
}

// add multiple values
func (l *singlyLinkedlist) AddMultipleValues() {
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

		l.Addvalue(val) //pass the value to insert int linked list
	}
}

// to remove values from linked list
func (l *singlyLinkedlist) RemoveSingleValue(value int) {
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
func (l *singlyLinkedlist) PrintAllForward() {
	newNode := l.head

	if newNode == nil {
		fmt.Println("Linked list is empty")
		return
	}
	fmt.Println("\nLinked list values in Froward order")
	for newNode != nil {
		fmt.Print(newNode.data, " ")
		newNode = newNode.next
	}
	fmt.Println()
}

// to print all values in backward
func (l *singlyLinkedlist) PrintAllBackward() {

	currentNode := l.head

	if currentNode == nil {
		fmt.Println("Linked list is empty")
		return
	}
	fmt.Println("\nLinked list values in Backward order")

	printAllBackwardHelper(currentNode)
	fmt.Println()
}

func printAllBackwardHelper(node *node) {

	if node == nil {
		return
	}

	printAllBackwardHelper(node.next)
	fmt.Print(node.data, " ")
}
