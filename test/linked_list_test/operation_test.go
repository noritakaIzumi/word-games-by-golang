package linked_list_test

import (
	"reflect"
	"testing"
	"word_games/pkg/linked_list"
)

func TestLinkedList_AddAtHead(t *testing.T) {
	type fields struct {
		Head *linked_list.ListNode
	}
	type args struct {
		val string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// NOTE: ノードの順番は「レシーバの中身をチェックするテスト」で確かめる
		// {name: "nil receiver causes error.", fields: nil, args: args{val: "a"}, wantErr: true},
		{name: "empty node.", fields: fields{Head: nil}, args: args{val: "a"}, wantErr: false},
		{name: "some nodes.", fields: fields{Head: &linked_list.ListNode{
			Val:  "b",
			Next: nil,
		}}, args: args{val: "a"}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rcv := &linked_list.LinkedList{
				Head: tt.fields.Head,
			}
			if err := rcv.AddAtHead(tt.args.val); (err != nil) != tt.wantErr {
				t.Errorf("AddAtHead() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestLinkedList_AddAtTail(t *testing.T) {
	type fields struct {
		Head *linked_list.ListNode
	}
	type args struct {
		val string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// NOTE: ノードの順番は「レシーバの中身をチェックするテスト」で確かめる
		// {name: "nil receiver causes error.", fields: nil, args: args{val: "a"}, wantErr: true},
		{name: "empty node.", fields: fields{Head: nil}, args: args{val: "a"}, wantErr: false},
		{name: "some nodes.", fields: fields{Head: &linked_list.ListNode{
			Val: "a",
			Next: &linked_list.ListNode{
				Val:  "b",
				Next: nil,
			},
		}}, args: args{val: "c"}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rcv := &linked_list.LinkedList{
				Head: tt.fields.Head,
			}
			if err := rcv.AddAtTail(tt.args.val); (err != nil) != tt.wantErr {
				t.Errorf("AddAtTail() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestLinkedList_DeleteAtHead(t *testing.T) {
	type fields struct {
		Head *linked_list.ListNode
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// NOTE: ノードの順番は「レシーバの中身をチェックするテスト」で確かめる
		// {name: "nil receiver causes error.", fields: nil, wantErr: true},
		{name: "empty node.", fields: fields{Head: nil}, wantErr: false},
		{name: "1 node.", fields: fields{Head: &linked_list.ListNode{
			Val:  "a",
			Next: nil,
		}}, wantErr: false},
		{name: "2 nodes.", fields: fields{Head: &linked_list.ListNode{
			Val: "a",
			Next: &linked_list.ListNode{
				Val:  "b",
				Next: nil,
			},
		}}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rcv := &linked_list.LinkedList{
				Head: tt.fields.Head,
			}
			if err := rcv.DeleteAtHead(); (err != nil) != tt.wantErr {
				t.Errorf("DeleteAtHead() error = %v, wantErr %v", err, tt.wantErr)
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

func TestLinkedList_MidCut(t *testing.T) {
	type fields struct {
		Head *linked_list.ListNode
	}
	tests := []struct {
		name       string
		fields     fields
		wantNewLst *linked_list.LinkedList
		wantErr    bool
	}{
		// NOTE: レシーバー側のノードの順番は「レシーバの中身をチェックするテスト」で確かめる
		// {name: "nil receiver causes error.", fields: nil, wantErr: true},
		{name: "no node.", fields: fields{Head: nil}, wantNewLst: &linked_list.LinkedList{Head: nil}, wantErr: false},
		{name: "1 node.", fields: fields{Head: &linked_list.ListNode{
			Val:  "a",
			Next: nil,
		}}, wantNewLst: &linked_list.LinkedList{Head: &linked_list.ListNode{
			Val:  "a",
			Next: nil,
		}}, wantErr: false},
		{name: "2 nodes.", fields: fields{Head: &linked_list.ListNode{
			Val: "a",
			Next: &linked_list.ListNode{
				Val:  "b",
				Next: nil,
			},
		}}, wantNewLst: &linked_list.LinkedList{Head: &linked_list.ListNode{
			Val:  "a",
			Next: nil,
		}}, wantErr: false},
		{name: "3 nodes.", fields: fields{Head: &linked_list.ListNode{
			Val: "a",
			Next: &linked_list.ListNode{
				Val: "b",
				Next: &linked_list.ListNode{
					Val:  "c",
					Next: nil,
				},
			},
		}}, wantNewLst: &linked_list.LinkedList{Head: &linked_list.ListNode{
			Val: "b",
			Next: &linked_list.ListNode{
				Val:  "a",
				Next: nil,
			},
		}}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rcv := &linked_list.LinkedList{
				Head: tt.fields.Head,
			}
			gotNewLst, err := rcv.MidCut()
			if (err != nil) != tt.wantErr {
				t.Errorf("MidCut() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotNewLst, tt.wantNewLst) {
				t.Errorf("MidCut() gotNewLst = %v, want %v", gotNewLst, tt.wantNewLst)
			}
		})
	}
}

func TestLinkedList_Tail(t *testing.T) {
	type fields struct {
		Head *linked_list.ListNode
	}
	tests := []struct {
		name    string
		fields  fields
		want    *linked_list.ListNode
		wantErr bool
	}{
		// {name: "nil receiver causes error.", fields: nil, wantErr: true},
		{name: "no node.", fields: fields{Head: nil}, want: nil, wantErr: false},
		{name: "1 node.", fields: fields{Head: &linked_list.ListNode{
			Val:  "a",
			Next: nil,
		}}, want: &linked_list.ListNode{
			Val:  "a",
			Next: nil,
		}, wantErr: false},
		{name: "some nodes.", fields: fields{Head: &linked_list.ListNode{
			Val: "a",
			Next: &linked_list.ListNode{
				Val: "b",
				Next: &linked_list.ListNode{
					Val:  "c",
					Next: nil,
				},
			},
		}}, want: &linked_list.ListNode{
			Val:  "c",
			Next: nil,
		}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rcv := &linked_list.LinkedList{
				Head: tt.fields.Head,
			}
			got, err := rcv.Tail()
			if (err != nil) != tt.wantErr {
				t.Errorf("Tail() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Tail() got = %v, want %v", got, tt.want)
			}
		})
	}
}

// ここからレシーバの中身をチェックするテスト

func TestLinkedList_receiverが正しいLinkedListを持っているかどうか(t *testing.T) {
	name := "Receiver has collect list."
	t.Run(name, func(t *testing.T) {
		rcv := &linked_list.LinkedList{
			Head: nil,
		}
		_ = rcv.AddAtTail("a")
		_ = rcv.AddAtHead("b")
		_ = rcv.AddAtTail("c")
		_, _ = rcv.MidCut()
		_ = rcv.AddAtHead("d")
		_ = rcv.AddAtTail("e")
		_, _ = rcv.MidCut()
		_ = rcv.AddAtTail("f")
		_ = rcv.DeleteAtHead()
		wantRcv := &linked_list.LinkedList{
			Head: &linked_list.ListNode{
				Val: "e",
				Next: &linked_list.ListNode{
					Val:  "f",
					Next: nil,
				},
			},
		}
		if !reflect.DeepEqual(rcv, wantRcv) {
			t.Errorf("%s: rcv = %v, wantRcv %v", name, rcv, wantRcv)
		}
	})
}
