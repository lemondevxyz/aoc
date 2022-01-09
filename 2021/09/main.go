package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	input := `2199943210
3987894921
9856789892
8767896789
9899965678`

	input = strings.ReplaceAll(input, "\n", "")
	numbers := []int{}
	for i := 0; i < len(input); i++ {
		number, err := strconv.Atoi(string(input[i]))
		if err != nil {
			panic(err)
		}

		numbers = append(numbers, number)
	}

	risk := 0
	for i := 0; i < len(numbers); i++ {
		up := i - 10
		left := i - 1
		down := i + 10
		right := i + 1

		val := numbers[i]
		if up >= 0 && numbers[up] < val {
			continue
		}

		if left >= 0 && numbers[left] < val {
			continue
		}

		if down < len(input) && numbers[down] < val {
			continue
		}

		if right < len(input) && numbers[right] < val {
			continue
		}

		risk += val + 1
	}

	fmt.Println("risk =", risk)
}
