package main

import (
	"fmt"
	d "minesweeper-go/definitions"
	_ "strconv"
)

func main() {

	fmt.Println(d.Demo(), " player")

	board := d.Board{}
	board.Level = 10
	mines := make(d.Mines, 0)

	board.M, board.N = 10, 20
	//	var mines d.Mines //:= make(map[d.Coord]bool)

	board.Mat = make([][]int32, board.M)
	for i := uint32(0); i < board.M; i++ {
		board.Mat[i] = make([]int32, board.N)
	}
	for i := uint32(0); i < board.M; i++ {
		for j := uint32(0); j < board.N; j++ {
			board.Mat[i][j] = -1
		}
	}

	d.PlaceMines(&board, &mines)
	fmt.Printf("mines\n%v\n", mines)
	fmt.Printf("board\n%v\n", board)

	// a := d.Coord{2, 3}	//testing imported package definitions
	// fmt.Println(a)

}
