package clist

import (
	"fmt"
	"time"

	"github.com/nikhilnarayanan623/ds-linked-list/clist/interfaces"
)

type node struct {
	data   int
	idx    int
	status bool
	next   *node
}

type circularList struct {
	head   *node
	tail   *node
	Length int
}

func NewcircularList() interfaces.CircularLinkedListInterface {

	return &circularList{Length: 0}
}

// complexity
// O(1)TS
// time always constant because adding a new node to tail its not depend on input
// space always take 1 because alwatys allocate 1 memory for a new node is constant
func (c *circularList) Append(data int) {

	newNode := &node{data: data, status: true}

	if c.head == nil {
		c.Length = 1
		c.head = newNode
	} else {
		c.tail.next = newNode
	}
	c.tail = newNode
	c.tail.next = c.head
	//insert intex to last value

	c.tail.idx = c.Length
	c.Length++
}

// complexity
// O(n)TS
// time n, always depend on user input because loop is working according to user input
// space n, calling addValues function create new node so n times memmory needed so space is 'n'
func (c *circularList) AppendMultipleValues() {
	var (
		limit int
	)

	fmt.Print("How many numbers you need to add: ")
	fmt.Scan(&limit)

	for i := 1; i <= limit; i++ {
		//fmt.Print("Enter value: ")
		//fmt.Scan(&val)
		c.Append(i)
	}
	fmt.Println()
}

// complexity
// O(n)T , O(1)S
// time take according to the input that much time loop throgh the linked list
// no extra space needed we dont't create new space only using the same space
func (c *circularList) DisplayValues() {

	if c.head == nil {
		fmt.Println("There is no values in circular linked list")
		return
	}

	fmt.Print(c.head.data, " ")

	currentNode := c.head.next

	for currentNode != c.head && currentNode.next != currentNode {
		fmt.Print(currentNode.data, " ")

		currentNode = currentNode.next
	}
	fmt.Println()
}

// complexity
// O(n^2)T , O(1)S

// small game
// consider 100 people and given a gun
// first one kill next person and give the gun to next person who alive
// find the winner who is left at last
// in here we wont delete nodes only make node as false
func (c *circularList) FindWinner() int {

	killer := c.head
	victim := killer.next

	for killer != victim {

		if killer == c.head {
			fmt.Println("\n Round completed ")
		}

		victim.status = false // kill the victim

		fmt.Printf("%d th shoot--> %d person \n", killer.idx, victim.idx)
		time.Sleep(300 * time.Millisecond)

		/*
			//first approach

			killer = killer.next
			/// for the next kill
			for !killer.status { //find a living killer
				killer = killer.next
			}

			//find the next victim
			victim = killer.next

			for !victim.status { //to find living victim
				victim = victim.next
			}
		*/

		// second approach
		killer = killer.next
		victim = victim.next

		for !killer.status || !victim.status {

			if !killer.status {
				killer = killer.next
				victim = killer.next
			} else {
				victim = victim.next
			}
		}
		//
	}

	//approach third with delete the next node

	// for killer.next != killer {
	// 	fmt.Println(killer.idx)
	// 	killer.next = killer.next.next
	// 	killer = killer.next
	// }

	return killer.idx

}

// game to delete 3rd st node and return the value of last node
func (c *circularList) DeleteEveryThirdNode() int {

	first := c.head
	//second := c.head
	count := 1

	for first != first.next {

		if count == 2 { //then value need to delete
			first.next = first.next.next
			count = 0 //reset count
		}
		//move to next
		first = first.next
		count++

	}
	//set the last value as list head
	c.head = first
	return first.data
}
