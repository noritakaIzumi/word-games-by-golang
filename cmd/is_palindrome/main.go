package main

import (
	"fmt"

	"word_games/internal"
	"word_games/pkg/io"
	"word_games/pkg/message"
)

func main() {
	word := io.GetInput(message.PleaseEnterSomeWord)

	result := message.Empty
	if internal.IsPalindrome(word) {
		result = message.Palindrome
	} else {
		result = message.NotPalindrome
	}

	fmt.Printf(message.FmtPalindromeResult, word, result)
}
