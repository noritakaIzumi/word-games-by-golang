package linked_list_test

import (
	"reflect"
	"testing"
	"word_games/pkg/linked_list"
)

func TestConvertWordToLinkedList(t *testing.T) {
	type args struct {
		word linked_list.Word
	}
	tests := []struct {
		name string
		args args
		want *linked_list.LinkedList
	}{
		{name: "normal case.", args: args{word: linked_list.Word("apt")}, want: &linked_list.LinkedList{Head: &linked_list.ListNode{
			Val: "a",
			Next: &linked_list.ListNode{
				Val: "p",
				Next: &linked_list.ListNode{
					Val:  "t",
					Next: nil,
				},
			},
		}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := linked_list.ConvertWordToLinkedList(tt.args.word); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConvertWordToLinkedList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWord_IsPalindrome(t *testing.T) {
	tests := []struct {
		name    string
		rcv     linked_list.Word
		wantRet bool
	}{
		{name: "empty string.", rcv: linked_list.Word(""), wantRet: true},
		{name: "1 char.", rcv: linked_list.Word("x"), wantRet: true},
		{name: "odd chars: false case.", rcv: linked_list.Word("study"), wantRet: false},
		{name: "odd chars: true case.", rcv: linked_list.Word("level"), wantRet: true},
		{name: "even chars: false case.", rcv: linked_list.Word("palindrome"), wantRet: false},
		{name: "even chars: true case.", rcv: linked_list.Word("noon"), wantRet: true},
		{name: "マルチバイト文字。", rcv: linked_list.Word("トマト"), wantRet: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRet := tt.rcv.IsPalindrome(); gotRet != tt.wantRet {
				t.Errorf("IsPalindrome() = %v, want %v", gotRet, tt.wantRet)
			}
		})
	}
}
