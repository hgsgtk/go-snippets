package fizzbuzz_test

import (
	"testing"

	"github.com/hgsgtk/go-snippets/testing-codes/fizzbuzz"
)

func TestGetMsgTableDriven(t *testing.T) {
	tests := []struct {
		num  int
		want string
	}{
		{
			num:  15,
			want: "FizzBuzz",
		},
		{
			num:  5,
			want: "Buzz",
		},
		{
			num:  3,
			want: "Fizz",
		},
		{
			num:  1,
			want: "1",
		},
	}
	for _, tt := range tests {
		if got := fizzbuzz.GetMsg(tt.num); got != tt.want {
			t.Errorf("GetMsg(%d) = %s, want %s", tt.num, got, tt.want)
		}
	}
}

func TestGetMsgSubTest(t *testing.T) {
	tests := []struct {
		desc string
		num  int
		want string
	}{
		{
			desc: "divisible by 15",
			num:  15,
			want: "FizzBuzz",
		},
		{
			desc: "divisible by 5",
			num:  5,
			want: "Buzz",
		},
		{
			desc: "divisible by 3",
			num:  3,
			want: "Fizz",
		},
		{
			desc: "not divisible",
			num:  1,
			want: "1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			if got := fizzbuzz.GetMsg(tt.num); got != tt.want {
				t.Errorf("GetMsg(%d) = %s, want %s", tt.num, got, tt.want)
			}
		})
	}
}

func TestGetMsg2(t *testing.T) {
	var num int
	var want string

	num = 15
	want = "FizzBuzz"
	if got := fizzbuzz.GetMsg(num); got != want {
		t.Fatalf("GetMsg(%d) = %s, want %s", num, got, want)
	}

	num = 5
	want = "Buzz"
	if got := fizzbuzz.GetMsg(num); got != want {
		t.Fatalf("GetMsg(%d) = %s, want %s", num, got, want)
	}

	num = 3
	want = "Fizz"
	if got := fizzbuzz.GetMsg(num); got != want {
		t.Fatalf("GetMsg(%d) = %s, want %s", num, got, want)
	}

	num = 1
	want = "1"
	if got := fizzbuzz.GetMsg(num); got != want {
		t.Fatalf("GetMsg(%d) = %s, want %s", num, got, want)
	}
}
