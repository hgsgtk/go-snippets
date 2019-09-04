package main

import (
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	s := rand.Intn(6)
	switch s {
	case 5:
		println("大吉")
	case 4, 3:
		println("中吉")
	case 2, 1:
		println("小吉")
	case 0:
		println("凶")
	}
}
