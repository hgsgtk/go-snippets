package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"testing"
	"testing/quick"
)

const IntRange = 10000

// テスト用の値の幅をいじるために専用の型を定義
type testInt int

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

// IntRangeの範囲しか返さないようにする
func (i testInt) Generate(rand *rand.Rand, size int) reflect.Value {
	v := testInt(randInt(-1*IntRange, IntRange))
	return reflect.ValueOf(v)
}

func TestMultipleByThree(t *testing.T) {
	f := func(i testInt) bool {
		x := int(i)
		y := multipleByThree(x)
		return y/3 == x && y%3 == 0
	}

	// デフォルトだと100パターンしか試さないので、パターンを1000倍（100,000パターン）に増やしてみる
	c := &quick.Config{
		MaxCountScale: 1000,
	}
	if err := quick.Check(f, c); err != nil {
		t.Error(err)
	}
}

func multipleByThree(x int) int {
	fmt.Println(x)
	if x == 3 { // わざとらしいcorner case
		return 2
	}
	return x * 3
}
