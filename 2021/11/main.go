package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Matrix [][]int

func ParseMatrix(str string) (m Matrix) {
	split := strings.Split(str, "\n")
	for _, v := range split {
		index := len(m)
		m = append(m, []int{})
		for _, char := range v {
			num, err := strconv.Atoi(string(char))
			if err != nil {
				panic(fmt.Sprintf("bad input: %s", err.Error()))
			}
			m[index] = append(m[index], num)
		}
	}

	return
}

func (m Matrix) String() (str string) {
	for k, v := range m {
		if k != 0 {
			str += "\n"
		}
		for _, number := range v {
			str += strconv.Itoa(number)
		}
	}

	return
}

func (m Matrix) Add(y, x int) {
	if m[y][x] <= 9 && m[y][x] != 0 {
		m[y][x]++
	}
}

func (m Matrix) Flash(y, x int) {
	m[y][x] = 0
	exec := func(y, x int) {
		if y > -1 && y < len(m) && x > -1 && x < len(m[y]) {
			m.Add(y, x)
			if m[y][x] > 9 {
				m.Flash(y, x)
			}
		}
	}

	top := y - 1
	bottom := y + 1
	left := x - 1
	right := x + 1

	exec(top, x)
	exec(top, left)
	exec(top, right)
	exec(y, left)
	exec(y, right)
	exec(bottom, x)
	exec(bottom, left)
	exec(bottom, right)
}

func (m Matrix) Step() {
	for y := range m {
		for x := range m[y] {
			m[y][x]++
		}
	}

	for y := range m {
		for x := range m[y] {
			if m[y][x] >= 10 {
				m.Flash(y, x)
			}
		}
	}

}

func main() {
	str := `
5483143223
2745854711
5264556173
6141336146
6357385478
4167524645
2176841721
6882881134
4846848554
5283751526`
	/*
	   	str = `
	   11111
	   19991
	   19191
	   19991
	   11111`
	*/

	str = str[1:]

	m := ParseMatrix(str)
	fmt.Println("Before any steps:")
	fmt.Println(m)
	for i := 1; i <= 10; i++ {
		m.Step()
		fmt.Printf("\nAfter step %d:\n", i)
		fmt.Println(m)
	}
}
