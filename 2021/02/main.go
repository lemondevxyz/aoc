package main

import (
	"fmt"
	"strconv"
	"strings"
)

func process(str string, x, y int) (int, int) {

	split := strings.Split(str, " ")
	word := split[0]

	num, err := strconv.Atoi(split[1])
	if err != nil {
		panic(err)
	}

	switch word {
	case "forward":
		x += num
	case "down":
		y += num
	case "up":
		y = y - num
	}

	return x, y
}

func main() {
	orders := `forward 5
down 5
forward 8
up 3
down 8
forward 2`

	x, y := 0, 0
	for _, v := range strings.Split(orders, "\n") {
		x, y = process(v, x, y)
		fmt.Println(x, y)
	}
}
