package day13

import (
	"fmt"
	"math/big"
)

//var busLinesWithMinuteSkips = []int{1789, 37, 47, 1889}

//var busLinesWithMinuteSkips = []int{19, -1, -1, -1, -1, -1, -1, -1, -1, 41, -1, -1, -1, 37, -1, -1, -1, -1, -1, 787, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, 13, -1, -1, -1, -1, -1, -1, -1, -1, -1, 23, -1, -1, -1, -1, -1, 29, -1, 571, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, 17}

func bigBusMinuteInputToMap() map[int]*big.Int {
	res := map[int]*big.Int{}
	for i := range busLinesWithMinuteSkips {
		if busLinesWithMinuteSkips[i] == -1 {
			continue
		}
		val := big.NewInt(int64(i))
		res[busLinesWithMinuteSkips[i]] = val
	}
	return res
}

func bigCheckBusesAtTS(delaysByBusID map[int]*big.Int, startTS *big.Int) bool {
	for busID, desiredMod := range delaysByBusID {
		n := big.NewInt(0).Add(startTS, desiredMod)
		if n.Mod(n, big.NewInt(int64(busID))).Cmp(big.NewInt(0)) != 0 {
			return false
		}
	}
	return true
}

func bigFindLargestBusID() int {
	max := -1
	for _, busID := range busLinesWithMinuteSkips {
		if busID > max {
			max = busID
		}
	}
	return max
}

func BigConsecutiveBuses() {
	delaysByBusID := bigBusMinuteInputToMap()
	fmt.Printf("%#v\n", delaysByBusID)
	i := big.NewInt(1)
	one := big.NewInt(1)
	largestBusID := bigFindLargestBusID()
	multiplier := big.NewInt(1)
	firstBusTS := big.NewInt(1)
	zero := big.NewInt(0)
	for {
		multiplier.Mul(big.NewInt(int64(largestBusID)), i)
		firstBusTS.Sub(multiplier, delaysByBusID[largestBusID])
		i.Add(one, i)
		if big.NewInt(0).Set(firstBusTS).Mod(firstBusTS, big.NewInt(int64(busLinesWithMinuteSkips[0]))).Cmp(zero) != 0 {
			continue
		}
		if bigCheckBusesAtTS(delaysByBusID, firstBusTS) {
			fmt.Println("found the first minute: ", firstBusTS)
			return
		}
	}
}
