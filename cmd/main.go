package main

import (
	"flag"
	"word_games/cmd/is_palindrome"
)

func main() {
	flag.Parse()
	appName := flag.Arg(0)

	switch appName {
	case "is_palindrome":
		is_palindrome.Main()
	}
}
