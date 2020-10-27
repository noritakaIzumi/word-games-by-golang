package linked_list

type LinkedList struct {
	Head *ListNode
}

type ListNode struct {
	Val  Char
	Next *ListNode
}

type Word string
type Char = Word
