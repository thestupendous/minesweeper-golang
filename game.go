package main

import (
	d "github.com/thestupendous/minesweeper-golang/definitions"
	"fmt"
	_ "strconv"
)


func main() {

	 fmt.Println(d.Demo()," user")

	d.M,d.N = 10,20
	var mines d.Mines//:= make(map[d.Coord]bool)
	mines = d.Mines{}
	var board d.Board

	board = make([][]string,d.M)
	for i:=uint32(0);i<d.M;i++ {
		board[i] = make([]string,d.N)
	}
	for i:=uint32(0);i<d.M;i++{
		for j:=uint32(0);j<d.N;j++{
			board[i][j] = " "
		}
	}

	mines[d.Coord{3,4}] = true


	fmt.Println(mines)
	fmt.Println(board)

}