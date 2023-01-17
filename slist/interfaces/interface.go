package interfaces

// these function are basically need for linked list
type LinkedListInterface interface {
	Addvalue(data int)
	AddMultipleValues()
	RemoveSingleValue(value int)
	PrintAllForward()
	PrintAllBackward()
}
