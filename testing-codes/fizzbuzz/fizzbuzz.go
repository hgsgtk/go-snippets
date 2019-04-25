package fizzbuzz

import "strconv"

func GetMsg(num int) string {
	var res string
	switch {
	case num%15 == 0:
		res = "FizzBuzz"
	case num%5 == 0:
		res = "Buzz"
	case num%3 == 0:
		res = "Fizz"
	default:
		res = strconv.Itoa(num)
	}
	return res
}
