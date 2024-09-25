package board

type GameMap struct {
	size_x int
	size_y int
	board [][]string 
}

func setBoard(size_x int, size_y int) [][]string {
	board := make([][]string, 0)	

	for i := 0; i < size_x; i++ {
		line := make([]string, 0)
		for j := 0; j < size_y; j++ {
			if (j == 0 || j == (size_y - 1) || i == 0 || i == (size_x - 1)){
				line = append(line, "#")
			}else{
				line = append(line, " ")
			}
		}
		board = append(board, line)			
	}
	
	return board
}

func CreateMap(size_x int, size_y int) *GameMap {
	m := new(GameMap);
	m.board = setBoard(size_x, size_y)
	m.size_x = size_x
	m.size_y = size_y	
	return m
}
