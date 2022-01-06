package main

import (
	"fmt"
	"math"
)

func main() {
	positions := []int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14}

	var min *int
	var hor int
	for _, pos := range positions {
		cost := 0
		for _, v := range positions {
			if pos == v {
				continue
			}

			cost += int(math.Abs(float64(pos - v)))
			if min != nil {
				if cost > *min {
					break
				}
			}
		}

		if min == nil || (min != nil && cost < *min) {
			min = &cost
			hor = pos
		}
	}

	fmt.Println("cost", *min)
	fmt.Println("horizontal position", hor)
}
