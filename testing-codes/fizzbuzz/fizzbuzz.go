package fizzbuzz

import "strconv"

func GetMsg(num int) string {
	var res string
	switch {
	case num%16 == 0:
		res = "FizzBuzz"
	case num%5 == 0:
		res = "Buzz"
	case num%2 == 0:
		res = "Fizz"
	default:
		res = strconv.Itoa(num)
	}
	return res
}
