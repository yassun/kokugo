package kokugo

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func Print(s string) {

	lines := strings.Split(s, "\n")
	maxLen := getMaxLen(lines)

	wrokLines := make([]string, cap(lines))
	for i, line := range lines {
		spaceCnt := maxLen - utf8.RuneCountInString(line)
		for i := 0; i < spaceCnt; i++ {
			line = line + "　"
		}
		wrokLines[i] = line
	}

	wrokRunes := string2Runes(wrokLines)
	runes := rotate(wrokRunes)

	for i := range runes {
		fmt.Println(string(runes[i]))
	}

}

func getMaxLen(lines []string) (maxLen int) {
	for _, line := range lines {
		if lineLen := utf8.RuneCountInString(line); maxLen < lineLen {
			maxLen = lineLen
		}
	}
	return maxLen
}

func string2Runes(strings []string) [][]rune {
	var runes [][]rune
	for i, s := range strings {
		runes = append(runes, []rune{})
		for _, r := range s {
			runes[i] = append(runes[i], r)
		}
	}
	return runes
}

func rotate(stringRune [][]rune) [][]rune {
	runes := make([][]rune, len(stringRune[0]))
	for i := range runes {
		runes[i] = make([]rune, len(stringRune))
	}

	for y, _ := range stringRune {
		for x, _ := range stringRune[0] {
			runes[x][len(stringRune)-y-1] = stringRune[y][x]
		}
	}

	return runes
}
