package definitions

var M,N uint32			//dimensions row,columns
type Coord struct {
	X,Y uint32
}

type Mines map[Coord]bool


type Board [][]string
func (board Board) String() string {
	out:= "╔"
	for i:=0;i<len(board[0]);i++{
		out += "═"
	}
	out += "╗\n"
	for i:=0;i<len(board);i++{
		out += "║"
		for j:=0;j<len(board[i]);j++{
			out += board[i][j]
		}
		out +="║\n"
	}
	out+= "╚"
	for i:=0;i<len(board[0]);i++{
		out += "═"
	}
	out += "╝\n"
	return out
}