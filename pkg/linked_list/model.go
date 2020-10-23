package linked_list

func ConvertWordToLinkedList(word Word) *linkedList {
    lst := getEmptyLinkedList()
    for _, c := range word {
        _ = lst.addAtTail(string(c))
    }
    return lst
}

func (rcv *Word) IsPalindrome() (ret bool) {
    lst1 := ConvertWordToLinkedList(*rcv)
    lst2, _ := lst1.MidCut()
    ret, _ = lst1.IsEqualTo(lst2)
    return ret
}
