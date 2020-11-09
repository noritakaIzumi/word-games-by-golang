package linked_list

func ConvertWordToLinkedList(word Word) *LinkedList {
	lst := GetEmptyLinkedList()
	for _, c := range word {
		_ = lst.AddAtTail(Char(c))
	}
	return lst
}

func ConvertLinkedListToWord(lst *LinkedList) *Word {
	return nil
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
