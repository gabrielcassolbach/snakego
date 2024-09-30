package snake 

type Point struct {
	Px int
	Py int
}

type Snake struct {
	snake_len int
	point_x int 
	point_y int
	queue []Point
}

func NewSnake(x int, y int) *Snake {
	s := new(Snake)
	body := make([]Point, 0)
	s.queue = body
	s.point_x = x
	s.point_y = y
	s.snake_len = 1
	return s
}

func (snke *Snake) Enqueue(new_point Point) {
	snke.queue = append(snke.queue, new_point)
}

func (snke *Snake) SetX(x int) {
	snke.point_x = x
}

func (snke *Snake) SetY(y int) {
	snke.point_y = y
} 

func (snke *Snake) GetX() int {
	return snke.point_x
} 

func (snke *Snake) GetSnakeSize() int {
	return snke.snake_len
}

func (snke *Snake) GrowSnake() {
	snke.snake_len = snke.snake_len + 1
}

func (snke *Snake) GetY() int {
	return snke.point_y
}

func (snke *Snake) QTop() Point {
	snke.queue = snke.queue[1:]	
	return snke.queue[0]
}

func (snke *Snake) GetQueueSize() int {
	return len(snke.queue)	
}


