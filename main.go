package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var characterMap = map[rune][5]rune{
	'a': {'a', 'ā', 'á', 'ǎ', 'à'},
	'o': {'o', 'ō', 'ó', 'ǒ', 'ò'},
	'e': {'e', 'ē', 'é', 'ě', 'è'},
	'i': {'i', 'ī', 'í', 'ǐ', 'ì'},
	'u': {'u', 'ū', 'ú', 'ǔ', 'ù'},
	'ü': {'ü', 'ǖ', 'ǘ', 'ǚ', 'ǜ'},
}

var vowelOrder = map[rune]int{
	'a': 5,
	'o': 4,
	'e': 3,
	'i': 2,
	'u': 1,
	'ü': 0,
}

type syllable struct {
	letters string
	tone    int
}

func execute(raw string) string {
	syllables := parseToSyllables(raw)
	converted := ""
	for _, s := range syllables {
		converted += s.convert()
	}

	return converted
}

func parseToSyllables(raw string) []syllable {
	trimmed := strings.Trim(raw, "\n\t")
	re := regexp.MustCompile("[0-4]+")

	tones := re.FindAllString(trimmed, -1)
	strSlice := re.Split(trimmed, -1)

	syllables := []syllable{}
	for i, s := range strSlice {
		if s == "" {
			continue
		}
		tone, _ := strconv.Atoi(tones[i])
		syllables = append(syllables, syllable{letters: s, tone: tone})
	}
	return syllables
}

func (s *syllable) convert() string {
	if s.tone == 0 {
		return s.letters
	}

	mainVowel, mainVowelIndex := determineMainVowel(s.letters)
	return s.letters[:mainVowelIndex] + string(characterMap[mainVowel][s.tone]) + s.letters[mainVowelIndex+1:]
}

func determineMainVowel(letters string) (rune, int) {
	var maxVowelValue, maxVowelIndex int
	for i, r := range letters {
		if _, ok := vowelOrder[r]; !ok {
			continue
		}
		if vowelOrder[r] >= maxVowelValue {
			maxVowelValue = vowelOrder[r]
			maxVowelIndex = i
		}
	}

	return rune(letters[maxVowelIndex]), maxVowelIndex
}

func main() {
	fmt.Println(execute(os.Args[1]))
}
