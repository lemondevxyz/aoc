package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	txt := `00100
11110
10110
10111
10101
01111
00111
11100
10000
11001
00010
01010`

	reverse := []string{}
	split := strings.Split(txt, "\n")
	for i := 0; i < 5; i++ {
		val := ""
		for _, v := range split {
			val += string(v[i])
		}
		reverse = append(reverse, val)
	}

	gamma := ""
	epsilon := ""
	for i := 0; i < 5; i++ {
		c0, c1 := strings.Count(reverse[i], "0"), strings.Count(reverse[i], "1")
		if c0 > c1 {
			gamma += "0"
			epsilon += "1"
		} else {
			gamma += "1"
			epsilon += "0"
		}
	}
	gammadecimal, _ := strconv.ParseInt(gamma, 2, 64)
	fmt.Println("Gamma:", gamma, "//", "Gamma Decimal:", gammadecimal)
	epsilondecimal, _ := strconv.ParseInt(epsilon, 2, 64)
	fmt.Println("Epsilon:", epsilon, "//", "Epsilon Decimal:", epsilondecimal)
	fmt.Println("Multiplied together:", epsilondecimal*gammadecimal)
}
