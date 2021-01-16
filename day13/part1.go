package day13

import "fmt"

const arriveTimestamp = 1003240

var busLines = []int{19, 41, 37, 787, 13, 23, 29, 571, 17}

func ShortestWaitBus() {
	minWait := 1000000
	winningBus := 0
	for _, busLine := range busLines {
		wait := ((1 + (arriveTimestamp / busLine)) * busLine) - arriveTimestamp
		if wait < minWait {
			minWait = wait
			winningBus = busLine
		}
	}
	fmt.Println("bus line: ", winningBus, ", wait time: ", minWait)
	fmt.Println("result: ", winningBus*minWait)
}
