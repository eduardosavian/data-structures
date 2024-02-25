package structs

type Node struct {
	value *Node
	next int
	prev int
}

type LinkedList struct {
	head *Node
	tail *Node
}