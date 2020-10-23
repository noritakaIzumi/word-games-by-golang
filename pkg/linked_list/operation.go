package linked_list

import (
	"errors"

	"word_games/pkg/message"
)

// addAtHead はリストの先頭にノードを追加します。
//
// リストが nil の場合、エラーを返します。
func (rcv *linkedList) addAtHead(val string) error {
	if rcv == nil {
		return errors.New(message.Nil)
	}
	rcv.head = &listNode{
		val:  val,
		next: rcv.head,
	}
	return nil
}

// addAtTail はリストの末尾にノードを追加します。
//
// リストが nil の場合、エラーを返します。
func (rcv *linkedList) addAtTail(val string) error {
	if rcv == nil {
		return errors.New(message.Nil)
	}
	// ノードがない場合、先頭にノードを追加しても末尾にノードを追加しても結果は不変。
	if rcv.head == nil {
		_ = rcv.addAtHead(val)
		return nil
	}
	rcv.tail().next = &listNode{
		val:  val,
		next: nil,
	}
	return nil
}

// tail はリストの末尾にあるノードを取得します。
func (rcv *linkedList) tail() *listNode {
	cur := rcv.head
	for cur.next != nil {
		cur = cur.next
	}
	return cur
}

// deleteAtHead はリストの先頭にあるノードを削除します。
//
// リストにノードがない場合、何もしません。
// リストが nil の場合、エラーを返します。
func (rcv *linkedList) deleteAtHead() error {
	if rcv == nil {
		return errors.New(message.Nil)
	}
	if rcv.head == nil {
		// 何もしない
	} else if rcv.head.next == nil {
		rcv.head = nil
	} else {
		rcv.head = rcv.head.next
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
func (rcv *linkedList) MidCut() (newLst *linkedList, err error) {
	if rcv == nil {
		return nil, errors.New(message.Nil)
	}
	cur := rcv.head
	newLst = getEmptyLinkedList()
	for cur != nil {
		_ = newLst.addAtHead(rcv.head.val)
		if cur.next == nil {
			break
		}
		cur = cur.next.next
		_ = rcv.deleteAtHead()
	}
	return newLst, nil
}

// IsEqualTo は 2 つのリストが持つ値の順序が一致しているかどうかを判断します。
// 一致していれば true 、一致していなければ false を返します。
//
// ノードの個数が異なる 2 つのリストは、前方のノードの値によらず一致していないとみなされます。
// ノードがない 2 つのリストは一致しているとみなされます。
//
// 一方のノードが nil の場合、エラーを返します。
func (rcv *linkedList) IsEqualTo(tgt *linkedList) (ret bool, err error) {
	if rcv == nil || tgt == nil {
		return false, errors.New(message.Nil)
	}
	for rcv.head != nil && tgt.head != nil {
		if rcv.head.val != tgt.head.val {
			return false, nil
		}
		_ = rcv.deleteAtHead()
		_ = tgt.deleteAtHead()
	}
	if rcv.head == nil && tgt.head == nil {
		return true, nil
	}
	return false, nil
}
