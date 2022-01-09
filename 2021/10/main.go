package main

import (
	"fmt"
	"strings"
)

func errorScore(c rune) int {
	switch c {
	case ')':
		return 3
	case ']':
		return 57
	case '}':
		return 1197
	case '>':
		return 25137
	}
	return 0
}

func parse(str string) (rune, rune) {
	chunkOpenings := "<{[("
	chunkClosers := ">}])"

	expect := []rune{}
loop:
	for _, v := range str {
		for k, open := range chunkOpenings {
			if v == open {
				expect = append(expect, rune(chunkClosers[k]))
				continue loop
			}
		}
		for _, closer := range chunkClosers {
			if v == closer {
				if expect[len(expect)-1] == closer {
					expect = expect[:len(expect)-1]
				} else {
					return expect[len(expect)-1], closer
				}
			}
		}
	}

	return 0, 0
}

func main() {
	input := `{([(<{}[<>[]}>{[]{[(<()>
[[<[([]))<([[{}[[()]]]
[{[{({}]{}}([{[{{{}}([]
[<(<(<(<{}))><([]([]()
<{([([[(<>()){}]>(<<{{`

	total := 0
	for _, v := range strings.Split(input, "\n") {
		expected, found := parse(v)
		fmt.Printf("%s: ", v)
		if expected == 0 {
			fmt.Printf("success")
		} else {
			fmt.Printf("Expected %c, but found %c instead.", expected, found)
			total += errorScore(found)
		}
		fmt.Println()
	}

	fmt.Printf("\nTotal Score: %d\n", total)
}
