package snake 

type Snake struct {
	body []string
}

func NewSnake() *Snake {
	s := new(Snake)
	body := make([]string, 0)
	body = append(body, "*")
	s.body = body
	return s
}

func (snke *Snake) GrowSnake() {
	snke.body = append(snke.body, "*")
} 

