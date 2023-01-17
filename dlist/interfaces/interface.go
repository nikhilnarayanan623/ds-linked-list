package interfaces

type DoublyLinkedListInterface interface {
	Addvalue(data int)                            //add a new value to linked list
	AddMultipleValues()                           //add multiple values to linked list
	RemoveSingleValue(value int)                  //remove a single value to linked list
	PrintAllForward()                             //print all values in the list in a forward direction
	PrintAllBackward()                            //print all values in the list in a backward direction
	InsertDataAfterAValue(checkVal int, data int) //data need to insert after a specific value
}
