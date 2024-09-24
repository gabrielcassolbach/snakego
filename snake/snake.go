package snake 

import "fmt"

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

func (snke *Snake) PrintSnake() {
	for i := 0; i < len(snke.body); i++{
		fmt.Print(snke.body[i])
	}
}


