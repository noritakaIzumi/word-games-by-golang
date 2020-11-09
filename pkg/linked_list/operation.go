package linked_list

import (
	"errors"

	"word_games/pkg/message"
)

// AddAtHead はリストの先頭にノードを追加します。
//
// リストが nil の場合、エラーを返します。
func (rcv *LinkedList) AddAtHead(val Char) error {
	if rcv == nil {
		return errors.New(message.Nil)
	}
	rcv.Head = &ListNode{
		Val:  val,
		Next: rcv.Head,
	}
	return nil
}

// AddAtTail はリストの末尾にノードを追加します。
//
// リストが nil の場合、エラーを返します。
func (rcv *LinkedList) AddAtTail(val Char) error {
	if rcv == nil {
		return errors.New(message.Nil)
	}
	// NOTE: ノードがない場合、「次のノード」という概念が存在しない。
	//  先頭にノードを追加しても末尾にノードを追加しても結果は不変。
	if rcv.Head == nil {
		_ = rcv.AddAtHead(val)
		return nil
	}
	tail, _ := rcv.Tail()
	tail.Next = &ListNode{
		Val:  val,
		Next: nil,
	}
	return nil
}

// Tail はリストの末尾にあるノードを取得します。
func (rcv *LinkedList) Tail() (*ListNode, error) {
	if rcv == nil {
		return nil, errors.New(message.Nil)
	}
	cur := rcv.Head
	if cur == nil {
		return nil, nil
	}
	index := 0
	for cur.Next != nil {
		index++
		cur = cur.Next
	}
	// NOTE: for 文についての C1 カバレッジを測定するため。
	if index == 0 {
		// for 文の中に入らなかった
	}
	return cur, nil
}

// DeleteAtHead はリストの先頭にあるノードを削除します。
//
// リストにノードがない場合、何もしません。
// リストが nil の場合、エラーを返します。
func (rcv *LinkedList) DeleteAtHead() error {
	if rcv == nil {
		return errors.New(message.Nil)
	}
	if rcv.Head == nil {
		// 何もしない
	} else if rcv.Head.Next == nil {
		rcv.Head = nil
	} else {
		rcv.Head = rcv.Head.Next
	}
	return nil
}

// MidCut はリストを真ん中の位置で切り取り、先頭から半分までのリストを逆順にして返します。
// 元のリストは真ん中から末尾までのリストになります。
// ノード数が奇数の場合、中央のノードは切り取ったリストと切り取られたリスト両方に残ります。
//
// 例えば
// l -> i -> s -> t というリストを切ると i -> l というリストを返し、元のリストは s -> t になります。
// a -> p -> p -> l -> e というリストを切ると p -> p -> a というリストを返し、元のリストは p -> l -> e になります。
// x というリストを切ると x と x という 2 つのリストに分かれます。
//
// リストにノードがない場合、エラーを返します。
func (rcv *LinkedList) MidCut() (newLst *LinkedList, err error) {
	if rcv == nil {
		return nil, errors.New(message.Nil)
	}
	cur := rcv.Head
	newLst = GetEmptyLinkedList()
	count := 0
	for cur != nil {
		count++
		_ = newLst.AddAtHead(rcv.Head.Val)
		if cur.Next == nil {
			break
		}
		cur = cur.Next.Next
		_ = rcv.DeleteAtHead()
	}
	// NOTE: for 文についての C1 カバレッジを測定するため。
	if count == 0 {
		// for 文の中に入らなかった
	}
	return newLst, nil
}
