package day15

import "fmt"

//var elfGameInput = []int{0, 3, 6}
var elfGameInput = []int{2, 0, 6, 12, 1, 3}

func NthElfNumber(upto int) {
	ts := 0
	lastSaidAt := map[int]int{}
	for ; ts < len(elfGameInput)-1; ts++ {
		lastSaidAt[elfGameInput[ts]] = ts + 1
	}
	last := elfGameInput[len(elfGameInput)-1]
	for ; ts < upto-1; ts++ {
		//fmt.Println("last=", last, " all=", lastSaidAt)
		if when, ok := lastSaidAt[last]; ok {
			lastSaidAt[last] = ts + 1
			last = ts - when + 1
		} else {
			lastSaidAt[last] = ts + 1
			last = 0
		}
	}

	fmt.Println("last number is", last)
}
