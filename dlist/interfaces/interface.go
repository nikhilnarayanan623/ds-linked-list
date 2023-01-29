package interfaces

type DoublyLinkedListInterface interface {
	Append(data int)                           // insert a new value into linked list at the end
	Prepend(data int)                          // insert a new value into the starting of the linked list
	AppendMultipleValues()                     // insert multiple values from user at the end
	DeleteValue(value int)                     // delete a specific value from the linked list
	DisplayOrder()                             // display all values in the linked list in forward order
	DisplayReverse()                           // display all values in the linked list in reverse order
	InsertAfterAValue(checkVal int, data int)  // insert a new values after a specific values
	InsertBeforeAValue(checkVal int, data int) // insert a new value before a specific value
	ArrayToSlist(array []int)                  //insert an array values to a linked list
	RemoveDuplicates()                         // to remove duplicate values if its in a sorted way
}
