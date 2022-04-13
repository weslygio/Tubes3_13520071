package main

import (
	"fmt"

	"github.com/weslygio/Tubes3_13520071/src/pkg/strmatcher"
)

func main() {
	text := "abacaabadcabacabaabb"
	pattern := "abacab"
	fmt.Println(strmatcher.Kmp(text, pattern))
	fmt.Println(strmatcher.Bm(text, pattern))

}
