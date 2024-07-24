package main

import (
	"aho-corasick/ahocorasick"
	"fmt"
)

func main() {
	keywords := []string{
		"he",
		"she",
		"his",
		"hers",
	}
	text := "ahishers"
	ans := map[string][]int{
		"his":  {1},
		"she":  {3},
		"he":   {4},
		"hers": {4},
	}
	ac := ahocorasick.NewACAutomaton(keywords)
	fmt.Println(ac)
	res := ac.Search(text)
	for k, v := range res {
		if len(v) != len(ans[k]) {
			panic("length not equal")
		}
		for i := range v {
			if v[i] != ans[k][i] {
				panic("value not equal")
			}
		}
	}
	fmt.Println("ok")
}
