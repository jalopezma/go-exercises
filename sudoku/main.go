package main

import (
	"fmt"

	"github.com/jalopezma/go-exercises/sudoku/sudoku"
)

func main() {
	sudoku := sudoku.Sudoku{}
	sudoku.Create()
	sudoku.Solve()
	fmt.Println("\nSolution\n", sudoku.String())
}
