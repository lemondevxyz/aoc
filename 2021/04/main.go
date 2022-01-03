package main

import "fmt"

type BoardItem struct {
	Value  int
	Marked bool
}

type Board [5][5]BoardItem

func (b Board) String() (str string) {
	for i := 0; i < 5; i++ {
		if i != 0 {
			str += "\n"
		}

		for j := 0; j < 5; j++ {
			m := "U"
			if b[i][j].Marked {
				m = "M"
			}
			str += fmt.Sprintf("%s%02d ", m, b[i][j].Value)
		}
	}

	return str
}

func (b *Board) Sum() (res int) {
	for i := range b {
		for j := range b[i] {
			if !b[i][j].Marked {
				res += b[i][j].Value
			}
		}
	}

	return
}

func (b *Board) Mark(value int) {
	for i := range b {
		for j := range b[i] {
			if b[i][j].Value == value {
				b[i][j].Marked = true
			}
		}
	}
}

func (b *Board) Complete() bool {
	for i := 0; i < 5; i++ {
		mx := 0
		my := 0
		for j := 0; j < 5; j++ {
			if b[i][j].Marked {
				mx++
			}
			if b[j][i].Marked {
				my++
			}
		}

		if my == 5 || mx == 5 {
			return true
		}
	}

	return false
}

func NewBoard(x [5][5]int) *Board {
	b := Board{}
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			b[i][j].Value = x[i][j]
		}
	}

	return &b
}

func main() {
	b1 := NewBoard([5][5]int{
		{22, 13, 17, 11, 0},
		{8, 2, 23, 4, 24},
		{21, 9, 14, 16, 7},
		{6, 10, 3, 18, 5},
		{1, 12, 20, 15, 19},
	})

	b2 := NewBoard([5][5]int{
		{3, 15, 0, 2, 22},
		{9, 18, 13, 17, 5},
		{19, 8, 7, 25, 23},
		{20, 11, 10, 24, 4},
		{14, 21, 16, 12, 6},
	})

	b3 := NewBoard([5][5]int{
		{14, 21, 17, 24, 4},
		{10, 16, 15, 9, 19},
		{18, 8, 23, 26, 20},
		{22, 11, 13, 6, 5},
		{2, 0, 12, 3, 7},
	})

	steps := []int{7, 4, 9, 5, 11, 17, 23, 2, 0, 14, 21, 24, 10, 16, 13, 6, 15, 25, 12, 22, 18, 20, 8, 19, 3, 26, 1}

	var winner *Board
	i := 0

	for ; i < len(steps); i++ {
		b1.Mark(steps[i])
		b2.Mark(steps[i])
		b3.Mark(steps[i])

		if b1.Complete() {
			fmt.Println("board 1 finishes first")
			winner = b1
			break
		} else if b2.Complete() {
			fmt.Println("board 2 finishes first")
			winner = b2
			break
		} else if b3.Complete() {
			fmt.Println("board 3 finishes first")
			winner = b3
			break
		}
	}

	fmt.Println("\n" + winner.String() + "\n")
	fmt.Println(winner.Sum(), "*", steps[i], "=", steps[i]*winner.Sum())

}
