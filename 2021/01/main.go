package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	str := `199
200
208
210
200
207
240
269
260
263`
	var last *int
	for _, v := range strings.Split(str, "\n") {
		num, err := strconv.Atoi(v)

		if err != nil {
			panic(err)
		}

		if last == nil {
			fmt.Println(v, "(N/A - no previous measurement)")
			last = new(int)
		} else {
			if *last > num {
				fmt.Println(v, "(decreased)")
			} else {
				fmt.Println(v, "(increased)")
			}
		}
		*last = num
	}
}
