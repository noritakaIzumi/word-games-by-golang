package message

const (
	Empty                 = ""
	PleaseEnterSomeWord   = "Please enter some word: "
	PleaseEnterFirstWord  = "Please enter first word: "
	PleaseEnterSecondWord = "Please enter second word: "
)

// is_palindrome
const (
	Palindrome    = "is palindrome!"
	NotPalindrome = "is not palindrome..."

	// FmtPalindromeResult は回文ゲームの結果を表示します。
	// 以下のように使用するとよいでしょう。
	//  fmt.Printf(message.FmtPalindromeResult, "apple", "is not palindrome...")
	FmtPalindromeResult = "The word \"%s\" %s\n"
)

// get_common_ending
//goland:noinspection SpellCheckingInspection
const (
	// FmtWordsHaveCommonEnding は共通する語尾を表示するためのフォーマットです。共通する語尾がある場合はこちら。
	//  fmt.Printf(message.FmtWordsHaveCommonEnding, "translation", "conversation", "ation")
	FmtWordsHaveCommonEnding = "The word \"%s\" and the word \"%s\" have the common ending \"%s\".\n"
	// FmtWordsHaveNoCommonEndings は共通する語尾を表示するためのフォーマットです。共通する語尾がない場合はこちら。
	//  fmt.Printf(message.FmtWordsHaveNoCommonEndings, "personal", "computer")
	FmtWordsHaveNoCommonEndings = "The word \"%s\" and the word \"%s\" have no common endings.\n"
)
