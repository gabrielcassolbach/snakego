package io

import (
	"bufio"
	"os"
)

func GetPressedKey() string {
	reader := bufio.NewReader(os.Stdin)
	key, _ := reader.ReadString('\n')
	return key
}
