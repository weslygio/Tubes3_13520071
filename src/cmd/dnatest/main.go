package main

import (
	"fmt"
	"io/ioutil"

	"github.com/weslygio/Tubes3_13520071/src/pkg/dna"
	"github.com/weslygio/Tubes3_13520071/src/pkg/strmatcher"
)

func main() {
	text := fileToString("../../test/text.txt")
	pattern := fileToString("../../test/pattern.txt")
	fmt.Println(strmatcher.Kmp(text, pattern))
	fmt.Println(strmatcher.Bm(text, pattern))
	matched, time, disease := dna.ParsePrediction("01-Apr-2022 HIV")
	fmt.Println(matched)
	fmt.Println(time)
	fmt.Println(disease)
}

func fileToString(filepath string) string {
	fileBytes, err := ioutil.ReadFile(filepath)
	if err != nil {
		panic(err)
	}
	return string(fileBytes)
}
