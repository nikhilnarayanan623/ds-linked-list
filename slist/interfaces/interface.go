package interfaces

// these function are basically need for linked list
type LinkedListInterface interface {
	Append(data int)
	Prepend(data int)                          //to insert a new value to then end of linked list
	AppendMultipleValues()                     // to insert multiple values to the linked list, limit is take from user
	DeleteValue(value int)                     // delete a specific value from linked list
	InsertAfterAValue(checkVal int, data int)  // insert a new value after a specific value
	InsertBeforeAValue(checkVal int, data int) // insert a new value before a specific value
	DisplayOrder()                             // display all values in the linked list in order
	DisplayReverse()                           // display all values in the linked list in reverse oreder
	RemoveDuplicates()                         //if singly linked list in a sortd form
	ArrayToSlist(array []int)                  //insert an array values to a linked list
}
