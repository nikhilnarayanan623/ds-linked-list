package interfaces

type CircularLinkedListInterface interface {
	Append(data int)
	AppendMultipleValues()
	DisplayValues()
	FindWinner() int
	DeleteEveryThirdNode() int
}
