package main

import (
	"fmt"
	"log"
	"os"

	"github.com/romandkv/sudoku/pkg/parser"
	"github.com/romandkv/sudoku/pkg/solver"
)

func main() {
	var ok bool
	if len(os.Args) != 2 {
		log.Fatalln("No map provided")
	}
	sudokuMap, err := parser.GetSudokuMap(os.Args[1])
	if err != nil {
		log.Fatalln(err)
	}
	*sudokuMap, ok = solver.Solve(*sudokuMap)
	if !ok {
		log.Fatalln("Unsolvable map")
	}
	fmt.Println(ok)
	solver.PrintMap(*sudokuMap)
}
