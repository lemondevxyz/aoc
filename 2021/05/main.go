package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Diagram [10][10]int

func (d Diagram) String() (str string) {
	for i := 0; i < len(d); i++ {
		if i != 0 {
			str += "\n"
		}

		for j := 0; j < len(d[i]); j++ {
			if d[i][j] == 0 {
				str += "."
			} else {
				str += fmt.Sprintf("%d", d[i][j])
			}
		}
	}

	return
}

type Point struct {
	X int
	Y int
}

type FromTo struct {
	From Point
	To   Point
}

func tonum(str string) int {
	v, _ := strconv.Atoi(str)
	return v
}

func parse_input(str string) []FromTo {
	arr := []FromTo{}

	str = strings.ReplaceAll(str, " -> ", ",")
	for _, v := range strings.Split(str, "\n") {
		split := strings.Split(v, ",")

		x1, y1, x2, y2 := tonum(split[0]), tonum(split[1]), tonum(split[2]), tonum(split[3])

		if x1 == x2 || y1 == y2 {
			arr = append(arr, FromTo{
				Point{x1, y1},
				Point{x2, y2},
			})
		}
	}

	return arr
}

func main() {
	str := `0,9 -> 5,9
8,0 -> 0,8
9,4 -> 3,4
2,2 -> 2,1
7,0 -> 7,4
6,4 -> 2,0
0,9 -> 2,9
3,4 -> 1,4
0,0 -> 8,8
5,5 -> 8,2`

	arr := parse_input(str)
	diagram := Diagram{}
	for _, v := range arr {
		from := v.From
		to := v.To

		var start, end int
		xbool := false
		if from.X == to.X {
			xbool = true
			start, end = from.Y, to.Y
		} else if from.Y == to.Y {
			start, end = from.X, to.X
		}

		if start > end {
			start, end = end, start
		}

		for i := start; i <= end; i++ {
			if xbool {
				diagram[i][to.X]++
			} else {
				diagram[to.Y][i]++
			}
		}
	}

	fmt.Println(diagram)

	for y, v := range diagram {
		for x, num := range v {
			if num == 2 {
				fmt.Println(x, y)
			}
		}
	}
}
