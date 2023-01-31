package main

import (
	"fmt"

	"github.com/nikhilnarayanan623/linked-list/clist"
	"github.com/nikhilnarayanan623/linked-list/dlist"
	"github.com/nikhilnarayanan623/linked-list/slist"
)

func binarySearch(array []int, target int) bool {

	var start, end = 0, len(array) - 1

	var mid int

	for start <= end {

		mid = start + ((end - start) / 2)

		if target == array[mid] {
			return true
		}

		if target < array[mid] {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}

	return false

}


//all linked lists are create in packages

// main file
// func signly for chekc signly linked list test
// func doubly for chekc signly linked list test
// func circular for chekc signly linked list test
func main() {

	fmt.Println(binarySearch([]int{1, 2, 3, 4, 5, 6}, 10))
}

func singly() {
	var data, checkVal int

	//singly liniked list
	fmt.Println("singly linked list")
	sl := slist.NewSinglyLinkedList() //get new singly linked list
	//to add multiple values
	sl.AppendMultipleValues()

	sl.DisplayOrder()

	//insert a values in the list after a specific number
	fmt.Println("Insert value after a specific value")
	fmt.Print("Enter the number for check: ")
	fmt.Scan(&checkVal)
	fmt.Print("Enter the number you want to insert: ")
	fmt.Scan(&data)
	sl.InsertAfterAValue(checkVal, data)
	sl.DisplayOrder()

	//insert a values in the list before a specific number
	fmt.Println("Insert value before a specific value")
	fmt.Print("Enter the number for check: ")
	fmt.Scan(&checkVal)
	fmt.Print("Enter the number you want to insert: ")
	fmt.Scan(&data)
	sl.InsertBeforeAValue(checkVal, data)
	sl.DisplayOrder()

	//insert a node to the begining of list
	fmt.Print("Enter a value to inset begining of list: ")
	fmt.Scan(&data)
	sl.Prepend(data)
	sl.DisplayOrder()

	//delete a value from the list
	fmt.Print("Enter a value to delete from list: ")
	fmt.Scan(&data)
	sl.DeleteValue(data)
	sl.DisplayOrder()

	//delet duplicate from linked list if its in a sorted order
	fmt.Println("After removing duplicate from list if its in a sorted way")
	sl.RemoveDuplicates()
	sl.DisplayOrder()
}

func doubly() {
	var data, checkVal int

	//doubly liniked list
	fmt.Println("doubly linked list")
	dl := dlist.NewDoublyLinkedList() //get new singly linked list
	//to add multiple values
	dl.AppendMultipleValues()
	dl.DisplayOrder()

	//insert a values in the list after a specific number
	fmt.Println("Insert value after a specific value")
	fmt.Print("Enter the number for check: ")
	fmt.Scan(&checkVal)
	fmt.Print("Enter the number you want to insert: ")
	fmt.Scan(&data)
	dl.InsertAfterAValue(checkVal, data)
	dl.DisplayOrder()

	//insert a values in the list before a specific number
	fmt.Println("Insert value before a specific value")
	fmt.Print("Enter the number for check: ")
	fmt.Scan(&checkVal)
	fmt.Print("Enter the number you want to insert: ")
	fmt.Scan(&data)
	dl.InsertBeforeAValue(checkVal, data)
	dl.DisplayOrder()

	//insert a node to the begining of list
	fmt.Print("Enter a value to inset begining of list: ")
	fmt.Scan(&data)
	dl.Prepend(data)
	dl.DisplayOrder()

	//delete a value from the list
	fmt.Print("Enter a value to delete from list: ")
	fmt.Scan(&data)
	dl.DeleteValue(data)
	dl.DisplayOrder()

	//delet duplicate from linked list if its in a sorted order
	fmt.Println("\nAfter removing duplicate from list if its in a sorted way")
	dl.RemoveDuplicates()
	dl.DisplayOrder()
	dl.DisplayReverse()

}

func circular() {

	//circular linked list

	c1 := clist.NewcircularList()

	c1.AppendMultipleValues()

	// fmt.Println("Values in the circular linked list")
	// c1.DisplayValues()

	winner := c1.FindWinner()

	fmt.Println("winner is at ", winner, "th position")
}
