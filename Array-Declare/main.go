package main

import (
	"log"
)

func main() {
	// var example - slices
	var a []int
	log.Printf("a type=%T, a length=%d, a capacity=%d, a=%v, nil? %t\n", a, len(a), cap(a), a, a == nil)

	// new example - slices
	b := new([]int)
	log.Printf("b type=%T, b length=%d, b capacity=%d, b=%v, nil? %t\n", b, len(*b), cap(*b), b, b == nil)

	// make examples = slices
	c := make([]int, 5)
	log.Printf("c type=%T, c length=%d, c capacity=%d, c=%v, nil? %t\n", c, len(c), cap(c), c, c == nil)

	d := make([]int, 5, 100)
	log.Printf("d type=%T, d length=%d, d capacity=%d, d=%v, nil? %t\n", d, len(d), cap(d), d, d == nil)

	// var example - array
	var e [3]int
	//log.Printf("e type=%T, e length=%d, e capacity=%d, e=%v, nil? %t\n", e, len(e), cap(e), e, e == nil)
	log.Printf("e type=%T, e length=%d, e capacity=%d, e=%v\n", e, len(e), cap(e), e)

	// new example - array
	f := new([3]int)
	log.Printf("f type=%T, f length=%d, f capacity=%d, f=%v, nil? %t\n", f, len(f), cap(f), f, f == nil)

	// slice literals
	g := []int{1, 3, 6}
	log.Printf("g type=%T, g length=%d, g capacity=%d, g=%v, nil? %t\n", g, len(g), cap(g), g, g == nil)

	h := []int{}
	log.Printf("h type=%T, h length=%d, h capacity=%d, h=%v, nil? %t\n", h, len(h), cap(h), h, h == nil)

	i := []champion{
		{
			Name:    "Evelynn",
			Classes: []string{"Assassin"},
			Origins: []string{"Demon"},
			Cost:    3,
		},
		{
			Name:    "Vi",
			Classes: []string{"Brawler"},
			Origins: []string{"Hextech"},
			Cost:    3,
		},
	}
	log.Printf("i type=%T, i length=%d, i capacity=%d, i=%v, nil? %t\n", i, len(i), cap(i), i, i == nil)

	// array literals
	j := [3]int{1, 3, 5}
	log.Printf("j type=%T, j length=%d, j capacity=%d, j=%v\n", j, len(j), cap(j), j)

	k := [...]int{1, 3, 5, 7}
	log.Printf("k type=%T, k length=%d, k capacity=%d, k=%v\n", k, len(k), cap(k), k)
}
