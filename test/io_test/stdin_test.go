package io_test

import (
	"testing"
	"word_games/pkg/io"
)

func TestGetInput(t *testing.T) {
	type args struct {
		msg string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := io.GetInput(tt.args.msg); got != tt.want {
				t.Errorf("GetInput() = %v, want %v", got, tt.want)
			}
		})
	}
}
