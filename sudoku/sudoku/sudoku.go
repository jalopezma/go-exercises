package sudoku

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

const max = 9

// Sudoku type
type Sudoku struct {
	cells [max][max]int
}

// Create - Creates a finished sudoku and remove some cells to create the puzzle
func (s *Sudoku) Create() {
	rand.Seed(time.Now().UTC().UnixNano())
	s.Solve()
	fmt.Printf("\nSudoku created\n%v\n", s)
	s.removeCells()
	fmt.Printf("\nPuzzle created\n%v\n", s)
}

func (s *Sudoku) removeCells() {
	removed := 0
	i := 0
	for {
		for j := 0; j < max; j++ {
			if rand.Intn(2) == 1 && s.cells[i][j] != 0 {
				removed++
				s.cells[i][j] = 0
				//fmt.Printf("Remove [%v, %v] removed %v\n", i, j, removed)
			}
			if removed >= 64 { //64
				break
			}
		}
		i = (i + 1) % max
		if removed >= 64 { //64
			break
		}
	}
}

func (s *Sudoku) getEmptyCell() (x, y int) {
	x, y = -1, -1
	for i, row := range s.cells {
		for j, value := range row {
			if value == 0 {
				return i, j
			}
		}
	}
	return x, y
}

// Solve -
func (s *Sudoku) Solve() (solved bool) {
	solved = false
	x, y := s.getEmptyCell()
	// fmt.Printf("Empty cell [%v, %v]\n", x, y)
	if x == -1 || y == -1 {
		return true
	}
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	rand.Shuffle(len(numbers), func(i, j int) {
		numbers[i], numbers[j] = numbers[j], numbers[i]
	})
	for _, num := range numbers {
		if s.isSafe(x, y, num) {
			s.cells[x][y] = num
			// fmt.Printf("[%v, %v] - %v\n", x, y, num)
			if s.Solve() {
				solved = true
				break
			}
			s.cells[x][y] = 0
		}
	}
	return solved
}

// isSafe - Checks if a number can be placed in a cell without breaking the 3x3, row and col
func (s *Sudoku) isSafe(x, y, n int) bool {
	return s.isSafeCol(y, n) && s.isSafeRow(x, n) && s.isSafeSquare(x, y, n)
}

// in - Checks if an int is included on an array
func (s *Sudoku) in(n int, list []int) bool {
	in := false
	for _, number := range list {
		in = number == n
		if in {
			break
		}
	}
	return in
}

// isSafeSquare - Checks if the square of 3/3 is valid
func (s *Sudoku) isSafeSquare(i, j, n int) bool {
	xSquare := i / 3
	ySquare := j / 3
	numbers := s.getNumbersSquare(xCube, yCube)
	r := !s.in(n, numbers)
	if !r {
		//	fmt.Printf("\nisSafeSquare [%v, %v] %v %v\n", i, j, n, numbers)
	}
	return r
}

// getNumbersSquare - Get the number of the square for the given point (for x1, y1 = square 0, 0)
func (s *Sudoku) getNumbersSquare(xCube, yCube int) []int {
	numbers := []int{}
	for i := xSquare * 3; i < (xCube+1)*3; i++ {
		for j := ySquare * 3; j < (yCube+1)*3; j++ {
			n := s.cells[i][j]
			if n != 0 {
				numbers = append(numbers, s.cells[i][j])
			}
		}
	}
	return numbers
}

// isSafeRow - Checks if the N number can be placed in the row
func (s *Sudoku) isSafeRow(i, n int) bool {
	numbers := s.getNumbersRow(i)
	r := !s.in(n, numbers)
	if !r {
		//fmt.Printf("\nisSafeRow [%v] %v %v\n", i, n, numbers)
	}
	return r
}

// getNumbersRow - Gets the numbers placed on the row given
func (s *Sudoku) getNumbersRow(row int) []int {
	numbers := []int{}
	for j := 0; j < max; j++ {
		n := s.cells[row][j]
		if n != 0 {
			numbers = append(numbers, s.cells[row][j])
		}
	}
	return numbers
}

// isSafeCol - Check if the N number can be placed in that col
func (s *Sudoku) isSafeCol(j, n int) bool {
	numbers := s.getNumbersCol(j)
	r := !s.in(n, numbers)
	if !r {
		//fmt.Printf("\nisSafeCol [%v] %v %v\n", j, n, numbers)
	}
	return r
}

// getNumbersCol - Gets the numbers placed on the col given
func (s *Sudoku) getNumbersCol(col int) []int {
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
func (s *Sudoku) String() string {
	str := ""
	for i, row := range s.cells {
		if i%3 == 0 {
			str = fmt.Sprintf("%v -----------------------\n", str)
		}
		for j, col := range row {
			if j%3 == 0 {
				str = fmt.Sprintf("%v|", str)
				if j != len(row) {
					str = fmt.Sprintf("%v ", str)
				}
			}
			value := strconv.Itoa(col)
			if value == "0" {
				value = " "
			}
			str = fmt.Sprintf("%s%v ", str, value)
		}
		str = fmt.Sprintf("%s|\n", str)
	}
	str = fmt.Sprintf("%v -----------------------", str)
	return str
}

// InitromArray - Inits the sudoku from an array
func (s *Sudoku) InitFromArray(numbers []int) bool {
	if len(numbers) < 81 {
		return false
	}
	for x := 0; x < len(s.cells); x++ {
		for y := 0; y < len(s.cells[x]); y++ {
			s.cells[x][y] = numbers[0]
			numbers = numbers[1:]
		}
	}
	return true
}

// InitromString - Inits the sudoku from the string representation
func (s *Sudoku) InitromString(str string) bool {
	numbers := []int{}
	x := 0
	y := 0
	for i := 0; i < len(str); i++ {
		char := string(str[i])
		// Number of lines the sudoku representation has
		if y >= 13 {
			break
		}
		if char == "-" {
			continue
		} else if char == "\n" {
			y++
			continue
		} else if char == "|" {
			x = 0
			continue
		} else if x%2 == 0 {
			x++
			continue
		}

		var n int
		var err error
		if char == " " {
			n = 0
		} else {
			n, err = strconv.Atoi(char)
			if err != nil {
				n = 0
			} else {
			}
		}
		numbers = append(numbers, n)
		x++
	}

	success := false
	if len(numbers) == 81 {
		success = s.InitFromArray(numbers)
	}
	return success
}
