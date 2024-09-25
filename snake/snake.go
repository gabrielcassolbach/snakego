package snake 

type Snake struct {
	point_x int 
	point_y int
	body []string
}

func NewSnake() *Snake {
	s := new(Snake)
	body := make([]string, 0)
	body = append(body, "*")
	s.body = body
	s.point_x = 15
	s.point_y = 35
	return s
}

func (snke *Snake) GrowSnake() {
	snke.body = append(snke.body, "*")
}

func (snke *Snake) GetX() int {
	return snke.point_x
} 

func (snke *Snake) GetY() int {
	return snke.point_y
}
