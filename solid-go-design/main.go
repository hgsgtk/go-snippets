// Open / Closed Principle > How does this advice apply to a language written 21 years later?
// package main

// import "fmt"

// type A struct {
// 	year int
// }

// func (a A) Greet() { fmt.Println("Hello Golang", a.year) }

// type B struct {
// 	A
// }

// func (b B) Greet() { fmt.Println("Welcome to Golang", b.year) }

// func main() {
// 	var a A
// 	a.year = 2016
// 	var b B
// 	b.year = 2016

// 	a.Greet()
// 	b.Greet()
// }

// Open / Closed Principle > So embedding is a powerful tool which allows Go’s types to be open for extension.
// package main

// import "fmt"

// type Cat struct {
// 	Name string
// }

// func (c Cat) Legs() int { return 4 }

// // before
// func (c Cat) PrintLegs() {
// 	fmt.Printf("I have %d legs\n", c.Legs())
// }

// type OctoCat struct {
// 	Cat
// }

// func (o OctoCat) Legs() int { return 5 }

// func main() {
// 	var octo OctoCat
// 	fmt.Println(octo.Legs())
// 	octo.PrintLegs()
// }
