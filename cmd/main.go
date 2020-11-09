package main

import (
	"flag"
	"word_games/cmd/get_common_ending"
	"word_games/cmd/is_palindrome"
)

func main() {
	flag.Parse()
	appName := flag.Arg(0)

	switch appName {
	case "is_palindrome":
		is_palindrome.Main()
	case "get_common_ending":
		get_common_ending.Main()
	}
}
