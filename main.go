package main

import (
	"aho-corasick/ahocorasick"
	"fmt"
)

func main() {
	keywords := []string{
		"say",
		"she",
		"shr",
		"he",
		"her",
	}
	ac := ahocorasick.NewACAutomaton(keywords)
	ac.PrintTrie()
	ans := ac.Search("sherhsay")
	fmt.Println(ans)
}
