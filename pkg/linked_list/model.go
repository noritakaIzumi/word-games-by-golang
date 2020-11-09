package linked_list

func ConvertWordToLinkedList(word Word) *LinkedList {
	lst := GetEmptyLinkedList()
	for _, c := range word {
		_ = lst.AddAtTail(Char(c))
	}
	return lst
}

func ConvertLinkedListToWord(list *LinkedList) *Word {
	w := ""
	cur := list.Head
	for cur != nil {
		w += string(cur.Val)
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

func GetCommonEnding(word1 *Word, word2 *Word) *Word {
	count := 0
	w1 := ConvertWordToLinkedList(*word1)
	w2 := ConvertWordToLinkedList(*word2)
	cur1 := w1.Head
	cur2 := w2.Head
	for count < 2 {
		if cur1.Next != nil {
			cur1 = cur1.Next
		} else {
			cur1 = w2.Head
			count++
		}
		if cur2.Next != nil {
			cur2 = cur2.Next
		} else {
			cur2 = w1.Head
			count++
		}
	}
	ce := cur1
	for cur1 != nil && cur2 != nil {
		if cur1.Val != cur2.Val {
			ce = cur1.Next
		}
		cur1 = cur1.Next
		cur2 = cur2.Next
	}
	result := ConvertLinkedListToWord(&LinkedList{Head: ce})
	return result
}
