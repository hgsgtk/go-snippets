package main

import "fmt"

func MultipleByThree(x int) int {
	fmt.Println(x)
	if x == 3 { // わざとらしいcorner case
		return 2
	}
	return x * 3
}

func main() {
	// 本来であればintの幅を超える場合とかもチェックすべきだけど今回はサンプルなので割愛
	fmt.Println(MultipleByThree(-1023568895))
}
