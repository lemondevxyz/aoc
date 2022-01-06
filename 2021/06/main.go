package main

import (
	"fmt"
	"strconv"
	"strings"
)

func process(str string) (ret string) {
	split := strings.Split(str, ",")
	for k := range split {
		num, err := strconv.Atoi(split[k])
		if err != nil {
			panic(err)
		}

		num--
		if num == -1 {
			num = 6
			split = append(split, "8")
		}

		split[k] = strconv.Itoa(num)
	}

	return strings.Join(split, ",")
}

func main() {
	input := "3,4,3,1,2"
	fmt.Println("Initial state:", input)
	days := 80
	for i := 1; i <= days; i++ {
		input = process(input)
		digit := fmt.Sprintf("%d", i)
		if i < 10 {
			digit = " " + digit
		}

		if i == 1 {
			fmt.Printf("After %s day:  ", digit)
		} else {
			fmt.Printf("After %s days: ", digit)
		}

		fmt.Println(input)
	}
}
