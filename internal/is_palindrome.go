package internal

import (
	"word_games/pkg/linked_list"
)

func IsPalindrome(word string) bool {
	w := linked_list.Word(word)
	return w.IsPalindrome()
}
