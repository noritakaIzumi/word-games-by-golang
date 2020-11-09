package internal

import "word_games/pkg/linked_list"

func GetCommonEnding(word1 string, word2 string) (bool, string) {
	w1 := linked_list.Word(word1)
	w2 := linked_list.Word(word2)
	ce := linked_list.GetCommonEnding(w1, w2)
	return len(ce) > 0, string(ce)
}
