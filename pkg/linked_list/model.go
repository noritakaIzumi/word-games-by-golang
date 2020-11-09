package linked_list

func ConvertWordToLinkedList(word Word) *LinkedList {
	lst := GetEmptyLinkedList()
	for _, c := range word {
		_ = lst.AddAtTail(Char(c))
	}
	return lst
}

func ConvertLinkedListToWord(lst *LinkedList) *Word {
	w := ""
	cur := lst.Head
	for cur != nil {
		w += string(lst.Head.Val)
		cur = cur.Next
	}
	result := Word(w)
	return &result
}

func (rcv *Word) IsPalindrome() (ret bool) {
	lst1 := ConvertWordToLinkedList(*rcv)
	lst2, _ := lst1.MidCut()
	ret, _ = lst1.IsEqualTo(lst2)
	return ret
}

func GetCommonEnding(w1 *Word, w2 *Word) *Word {
	return nil
}
