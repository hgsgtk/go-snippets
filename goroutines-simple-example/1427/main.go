package main

import (
	"fmt"
	"runtime"
	"sync"
)

// 計算・元金均等
func calc(id, price int, interestRate float64, year int) {
	months := year * 12
	interest := 0

	for i := 0; i < months; i++ {
		balance := price * (months - i) / months
		interest += int(float64(balance) * interestRate / 12)
	}
	fmt.Printf("year=%d total=%d interest=%d id=%d\n", year, price+interest, interest, id)
}

func worker(id, price int, interestRate float64, years chan int, wg *sync.WaitGroup) {
	for year := range years {
		calc(id, price, interestRate, year)
		wg.Done()
	}
}

func main() {
	price := 40000000
	interestRate := 0.011
	years := make(chan int, 35)

	for i := 1; i < 36; i++ {
		years <- i
	}

	var wg sync.WaitGroup
	wg.Add(35)

	for i := 0; i < runtime.NumCPU(); i++ {
		go worker(i, price, interestRate, years, &wg)
	}
	close(years)
	wg.Wait()
}
