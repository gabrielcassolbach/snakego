package io

import (
	"os"
)

type Keyboard struct {
	key string
}

func NewKeyboard() *Keyboard {
	kb := new(Keyboard)
	kb.key = "";
	return kb
}

func (kboard *Keyboard) AwaitKeyLoop() {
	var b []byte = make([]byte, 1)
	for {
		os.Stdin.Read(b)
		kboard.SetKey(string(b))
	}	
}

func (kboard *Keyboard) SetKey(k string) {
	kboard.key = k
}

func (kboard *Keyboard) GetKey() string{
	return kboard.key
}


