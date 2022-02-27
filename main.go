package main

import (
	"fmt"
	"regexp"
)

var characterMap = map[rune][4]rune{
	'a': {'ā', 'á', 'ǎ', 'à'},
	'o': {'ō', 'ó', 'ǒ', 'ò'},
	'e': {'ē', 'é', 'ě', 'è'},
	'i': {'ī', 'í', 'ǐ', 'ì'},
	'u': {'ū', 'ú', 'ǔ', 'ù'},
	'ü': {'ǖ', 'ǘ', 'ǚ', 'ǜ'},
}

func parseToSyllables(raw string) []string {
	re := regexp.MustCompile("[1-4]+")
	fmt.Println(re.FindAllString("ni3hao3", -1))
	fmt.Println(re.FindAllStringIndex("ni3hao3", -1))
	return []string{}
}

func convertSyllable(syllable string) string {
	return syllable
}

func main() {
	_ = parseToSyllables("")
	_ = convertSyllable("")
	fmt.Println(characterMap)
}
