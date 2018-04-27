package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/jalopezma/go-exercises/sudoku/sudoku"
)

var file string

func init() {
	flag.StringVar(&file, "file", "", "Input file with sudoku")
	flag.Parse()
}

func main() {
	sudoku := sudoku.Sudoku{}
	var content string
	if file != "" {
		content = readFile(file)
		if sudoku.InitFromString(content) {
			sudoku.Solve()
			fmt.Printf("\nSolution\n%v", sudoku.String())
		}
	} else {
		sudoku.Create()
		sudoku.Solve()
		fmt.Printf("\nSolution\n%v", sudoku.String())
	}
}

func readFile(fileName string) string {
	fmt.Printf("file %q\n", fileName)
	buff := make([]byte, 1024)
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
		return string(buff)
	}
	defer file.Close()

	for {
		n, err := file.Read(buff)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("Err % bytes: %v\n", n, err)
		}
	}
	return string(buff)
}
