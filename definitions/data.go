package definitions

import "fmt"

type Board struct {
	M, N  uint32
	Level uint16
	/*-10 means mine at current location
	0 means no mine at current or any neighboring location
	1 means one mine in all 8 neighbors
	2 means two mines in all 8 neighbors
	3 means three mines in all 8 neighbors
	4 means four mines in all 8 neighbors
	5 means five mines in all 8 neighbors
	6 means six mines in all 8 neighbors
	7 means seven mines in all 8 neighbors
	8 means eight mines in all 8 neighbors
	*/
	Mat [][]int32
}

type Coord struct {
	X, Y uint32
}

type Mines map[Coord]bool

//type  Board [][]string

func (board Board) String() string {
	// out := "╔"
	// for i := 0; i < len(board.Mat[0]); i++ {
	// out += "═"
	// }
	// out += "╗\n"
	out := ""
	for i := 0; i < len(board.Mat); i++ {
		// out += "║"
		for j := 0; j < len(board.Mat[i]); j++ {
			out += fmt.Sprintf("%3v", board.Mat[i][j])
		}
		// out += "║\n"
		out += "\n"
	}
	// out += "╚"
	// for i := 0; i < len(board.Mat[0]); i++ {
	// out += "═"
	// }
	// out += "╝\n"
	return out
}
