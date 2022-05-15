package definitions

import (
	"fmt"
	"math/rand"
	"time"
)

func Demo() string {
	return "hiii"
}

func PlaceMines(board *Board, mines *Mines) {
	source := rand.NewSource(time.Now().UnixNano())
	randGen := rand.New(source)

	// a := randGen.Intn(90)
	// fmt.Println(a, board.M, board.N)

	//TODO add support for levels
	var g, h uint32
	var i uint16
	for i = 0; i < (*board).Level; i++ {
		g, h = uint32(randGen.Intn(int((*board).M))), uint32(randGen.Intn(int((*board).N)))
		fmt.Println("g,h", g, h)
		(*board).Mat[g][h] = -8
		(*mines)[Coord{g, h}] = true

	}

}
