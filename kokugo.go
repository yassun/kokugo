package kokugo

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func Print(s string) {
	matrix := newMatrix(s)
	runes := matrix.rotate()
	for i := range runes {
		fmt.Println(string(runes[i]))
	}
}

func newInputLine(s string) *inputLine {
	return &inputLine{
		lines: format(s),
	}
}

type inputLine struct {
	lines []string
}

func (in *inputLine) string2Runes() [][]rune {
	var runes [][]rune
	for i, s := range in.lines {
		runes = append(runes, []rune{})
		for _, r := range s {
			runes[i] = append(runes[i], r)
		}
	}
	return runes
}

func newMatrix(s string) *matrix {
	inputLine := newInputLine(s)
	return &matrix{
		runes: inputLine.string2Runes(),
	}
}

type matrix struct {
	runes [][]rune
}

func (m *matrix) rotate() [][]rune {

	runes := make([][]rune, len(m.runes[0]))
	for i := range runes {
		runes[i] = make([]rune, len(m.runes))
	}

	for y, _ := range m.runes {
		for x, _ := range m.runes[0] {
			runes[x][len(m.runes)-y-1] = m.runes[y][x]
		}
	}

	return runes
}

func format(s string) []string {
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

	return wrokLines
}

func getMaxLen(lines []string) (maxLen int) {
	for _, line := range lines {
		if lineLen := utf8.RuneCountInString(line); maxLen < lineLen {
			maxLen = lineLen
		}
	}
	return maxLen
}
