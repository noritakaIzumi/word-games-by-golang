package get_common_ending

import (
	"fmt"
	"word_games/internal"
	"word_games/pkg/io"
	"word_games/pkg/message"
)

func Main() {
	firstWord := io.GetInput(message.PleaseEnterFirstWord)
	secondWord := io.GetInput(message.PleaseEnterSecondWord)

	hasCe, ce := internal.GetCommonEnding(firstWord, secondWord)
	if hasCe {
		fmt.Printf(message.FmtWordsHaveCommonEnding, firstWord, secondWord, ce)
	} else {
		fmt.Printf(message.FmtWordsHaveNoCommonEndings, firstWord, secondWord)
	}
}
