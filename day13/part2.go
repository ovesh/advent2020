package day13

import "fmt"

//var busLinesWithMinuteSkips = []int{1789, 37, 47, 1889}

var busLinesWithMinuteSkips = []int{19, -1, -1, -1, -1, -1, -1, -1, -1, 41, -1, -1, -1, 37, -1, -1, -1, -1, -1, 787, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, 13, -1, -1, -1, -1, -1, -1, -1, -1, -1, 23, -1, -1, -1, -1, -1, 29, -1, 571, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, 17}

func busMinuteInputToMap() map[int]int {
	res := map[int]int{}
	for i := range busLinesWithMinuteSkips {
		if busLinesWithMinuteSkips[i] == -1 {
			continue
		}
		res[busLinesWithMinuteSkips[i]] = i
	}
	return res
}

func checkBusesAtTS(delaysByBusID map[int]int, startTS int) bool {
	for busID, desiredMod := range delaysByBusID {
		if (startTS+desiredMod)%busID != 0 {
			return false
		}
	}
	return true
}

func findLargestBusID() int {
	max := -1
	for _, busID := range busLinesWithMinuteSkips {
		if busID > max {
			max = busID
		}
	}
	return max
}

func ConsecutiveBuses() {
	delaysByBusID := busMinuteInputToMap()
	fmt.Printf("%#v\n", delaysByBusID)
	i := 1
	largestBusID := findLargestBusID()
	for {
		firstBusTS := (largestBusID * i) - delaysByBusID[largestBusID]
		i++
		if firstBusTS%busLinesWithMinuteSkips[0] != 0 {
			continue
		}
		if checkBusesAtTS(delaysByBusID, firstBusTS) {
			fmt.Println("found the first minute: ", firstBusTS)
			return
		}
	}
}
