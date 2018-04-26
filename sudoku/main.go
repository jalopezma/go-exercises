package main

import (
	"fmt"

	"github.com/jalopezma/go-exercises/sudoku/sudoku"
)

func main() {
	sudoku := sudoku.Sudoku{}
	sudoku.Init()
	fmt.Println(sudoku.String())
}
