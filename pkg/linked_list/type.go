package linked_list

type linkedList struct {
    head *listNode
}

type listNode struct {
    val  string
    next *listNode
}

type Word string
