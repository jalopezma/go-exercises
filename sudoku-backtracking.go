package main

import (
	"fmt"
	"math/rand"
)

const max = 9

type sudoku struct {
	cells [max][max]int
}

// Init - Initializes the sudoku
func (s *sudoku) Init() {
	for i, _ := range s.cells {
		numbers := [max]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
		rand.Shuffle(len(numbers), func(i, j int) {
			numbers[i], numbers[j] = numbers[j], numbers[i]
		})
		numberIndex := 0
		j := 0
		for {
			n := numbers[numberIndex]
			fmt.Printf("i %v j %v n %v nI %v\n", i, j, n, numberIndex)
			if s.isSafeCol(j, n) && s.isSafeRow(i, n) && s.isSafeCube(i, j, n) {
				s.cells[i][j] = n
				j++
			}
			numberIndex = (numberIndex + 1) % max
			if j == max {
				break
			}
		}
	}
}

func (s *sudoku) isSafeCube(i, j, n int) bool {
	return true
}

func (s *sudoku) isSafeRow(i, n int) bool {
	return true
}

func (s *sudoku) isSafeCol(j, n int) bool {
	return true
}

// String - String representation
func (s *sudoku) String() string {
	str := ""
	for _, row := range s.cells {
		str = fmt.Sprintf("%s\n%v", str, row)
	}
	return str
}

func main() {
	sudoku := sudoku{}
	sudoku.Init()
	fmt.Println(sudoku.String())
}
