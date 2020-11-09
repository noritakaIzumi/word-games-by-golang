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

func TestLinkedList_IsEqualTo(t *testing.T) {
	type fields struct {
		Head *linked_list.ListNode
	}
	type args struct {
		tgt *linked_list.LinkedList
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantRet bool
		wantErr bool
	}{
		// {name: "nil receiver causes error. 1", fields: nil, args: args{tgt: nil}, wantErr: true},
		{name: "nil receiver causes error. 2", fields: fields{Head: nil}, args: args{tgt: nil}, wantErr: true},
		{name: "2 empty lists.", fields: fields{Head: nil}, args: args{tgt: &linked_list.LinkedList{Head: nil}}, wantRet: true, wantErr: false},
		{name: "different number of nodes.", fields: fields{Head: nil},
			args: args{tgt: &linked_list.LinkedList{Head: &linked_list.ListNode{
				Val:  "a",
				Next: nil,
			}}}, wantRet: false, wantErr: false},
		{name: "same number of nodes, different values.", fields: fields{Head: &linked_list.ListNode{
			Val: "a",
			Next: &linked_list.ListNode{
				Val:  "b",
				Next: nil,
			},
		}}, args: args{tgt: &linked_list.LinkedList{Head: &linked_list.ListNode{
			Val: "a",
			Next: &linked_list.ListNode{
				Val:  "c",
				Next: nil,
			},
		}}}, wantRet: false, wantErr: false},
		{name: "same number of nodes, same values.", fields: fields{Head: &linked_list.ListNode{
			Val: "a",
			Next: &linked_list.ListNode{
				Val:  "b",
				Next: nil,
			},
		}}, args: args{tgt: &linked_list.LinkedList{Head: &linked_list.ListNode{
			Val: "a",
			Next: &linked_list.ListNode{
				Val:  "b",
				Next: nil,
			},
		}}}, wantRet: true, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rcv := &linked_list.LinkedList{
				Head: tt.fields.Head,
			}
			gotRet, err := rcv.IsEqualTo(tt.args.tgt)
			if (err != nil) != tt.wantErr {
				t.Errorf("IsEqualTo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotRet != tt.wantRet {
				t.Errorf("IsEqualTo() gotRet = %v, want %v", gotRet, tt.wantRet)
			}
		})
	}
}

func TestConvertLinkedListToWord(t *testing.T) {
	type args struct {
		list *linked_list.LinkedList
	}
	tests := []struct {
		name    string
		args    args
		want    linked_list.Word
		wantErr bool
	}{
		{name: "empty string", args: args{list: &linked_list.LinkedList{Head: nil}}, want: "", wantErr: false},
		{name: "empty string", args: args{list: &linked_list.LinkedList{Head: &linked_list.ListNode{
			Val: "a",
			Next: &linked_list.ListNode{
				Val: "b",
				Next: &linked_list.ListNode{
					Val:  "c",
					Next: nil,
				},
			},
		}}}, want: linked_list.Word("abc"), wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := linked_list.ConvertLinkedListToWord(tt.args.list)
			if (err != nil) != tt.wantErr {
				t.Errorf("ConvertLinkedListToWord() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConvertLinkedListToWord() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetCommonEnding(t *testing.T) {
	type args struct {
		word1 linked_list.Word
		word2 linked_list.Word
	}
	tests := []struct {
		name string
		args args
		want linked_list.Word
	}{
		{name: "1 empty string", args: args{
			word1: "",
			word2: "apple",
		}, want: ""},
		{name: "same length strings & fragmentary common strings", args: args{
			word1: "aisle",
			word2: "waste",
		}, want: "e"},
		{name: "different length strings & continuous common strings", args: args{
			word1: "studying",
			word2: "developing",
		}, want: "ing"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := linked_list.GetCommonEnding(tt.args.word1, tt.args.word2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCommonEnding() = %v, want %v", got, tt.want)
			}
		})
	}
}
