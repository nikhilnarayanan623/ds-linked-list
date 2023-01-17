package main

import (
	"fmt"

	"github.com/nikhilnarayanan623/linkedlists/dlist"
	"github.com/nikhilnarayanan623/linkedlists/slist"
)

func main() {

	var data, checkVal int

	//singly liniked list
	fmt.Println("singly linked list")
	sl := slist.NewSinglyLinkedList() //get value from user
	sl.Addvalue(3)                    //to add a single value
	sl.AddMultipleValues()

	//doubly linked list
	fmt.Println("doubly linked list")

	dl := dlist.NewDoublyLinkedList()

	//get values from user
	dl.AddMultipleValues()

	//show the list
	dl.PrintAllForward()

	//insert a values in the list after a specific number
	fmt.Print("which number after you need to insert new value: ")
	fmt.Scan(&checkVal)
	fmt.Print("Enter the number you want to insert: ")
	fmt.Scan(&data)

	dl.InsertDataAfterAValue(checkVal, data)

	dl.PrintAllForward()
}
