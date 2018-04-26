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
		numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
		rand.Shuffle(len(numbers), func(i, j int) {
			numbers[i], numbers[j] = numbers[j], numbers[i]
		})
		numberIndex := 0
		j := 0
		tries := 0
		for {
			n := numbers[numberIndex]
			fmt.Printf("Shiffle: %v selected %v\n", numbers, n)
			fmt.Printf("i %v j %v n %v nI %v\n", i, j, n, numberIndex)
			safe := s.isSafeCol(j, n) && s.isSafeRow(i, n) && s.isSafeCube(i, j, n)
			if safe {
				s.cells[i][j] = n
				numbers = append(numbers[:numberIndex], numbers[numberIndex+1:]...)
				if len(numbers) > 0 {
					numberIndex = numberIndex % len(numbers)
				}
				j++
			} else {
				fmt.Println(s.String())
			}
			if !safe {
				numberIndex = (numberIndex + 1) % len(numbers)
				tries++
			}
			if tries > 40 {
				fmt.Println("tries > 40")
				break
			}
			if j == max {
				break
			}
		}
	}
}

func (s *sudoku) in(n int, list []int) bool {
	in := false
	for _, number := range list {
		in = number == n
		if in {
			break
		}
	}
	return in
}

func (s *sudoku) isSafeCube(i, j, n int) bool {
	xCube := i / 3
	yCube := j / 3
	numbers := s.getNumbersCube(xCube, yCube)
	r := !s.in(n, numbers)
	if !r {
		fmt.Printf("\nisSafeCube [%v, %v] %v %v\n", i, j, n, numbers)
	}
	return r
}

func (s *sudoku) getNumbersCube(xCube, yCube int) []int {
	numbers := []int{}
	for i := xCube * 3; i < (xCube+1)*3; i++ {
		for j := yCube * 3; j < (yCube+1)*3; j++ {
			n := s.cells[i][j]
			if n != 0 {
				numbers = append(numbers, s.cells[i][j])
			}
		}
	}
	return numbers
}

func (s *sudoku) isSafeRow(i, n int) bool {
	numbers := s.getNumbersRow(i)
	r := !s.in(n, numbers)
	if !r {
		fmt.Printf("\nisSafeRow [%v] %v %v\n", i, n, numbers)
	}
	return r
}

func (s *sudoku) getNumbersRow(row int) []int {
	numbers := []int{}
	for j := 0; j < max; j++ {
		n := s.cells[row][j]
		if n != 0 {
			numbers = append(numbers, s.cells[row][j])
		}
	}
	return numbers
}

func (s *sudoku) isSafeCol(j, n int) bool {
	numbers := s.getNumbersCol(j)
	r := !s.in(n, numbers)
	if !r {
		fmt.Printf("\nisSafeCol [%v] %v %v\n", j, n, numbers)
	}
	return r
}

func (s *sudoku) getNumbersCol(col int) []int {
	numbers := []int{}
	for i := 0; i < max; i++ {
		n := s.cells[i][col]
		if n != 0 {
			numbers = append(numbers, s.cells[i][col])
		}
	}
	return numbers
}

// String - String representation
func (s *sudoku) String() string {
	str := ""
	for i, row := range s.cells {
		if i%3 == 0 {
			str = fmt.Sprintf("%v------------------------------\n", str)
		}
		for j, col := range row {
			if j%3 == 0 {
				str = fmt.Sprintf("%v|", str)
			}
			str = fmt.Sprintf("%s %v ", str, col)
		}
		str = fmt.Sprintf("%s|\n", str)
	}
	str = fmt.Sprintf("%v------------------------------", str)
	return str
}

func main() {
	sudoku := sudoku{}
	sudoku.Init()
	fmt.Println(sudoku.String())
}
