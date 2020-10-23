package message

const (
	Empty               = ""
	PleaseEnterSomeWord = "Please enter some word: "
)

const (
	Palindrome    = "is palindrome!"
	NotPalindrome = "is not palindrome..."

	// FmtPalindromeResult は回文ゲームの結果を表示します。
	// 以下のように使用するとよいでしょう。
	//  fmt.Printf(message.FmtPalindromeResult, "apple", "is not palindrome...")
	FmtPalindromeResult = "The word \"%s\" %s\n"
)
