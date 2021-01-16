package day10

import (
	"fmt"
	"sort"
)

var joltageAdapters = []int{
	56,
	139,
	42,
	28,
	3,
	87,
	142,
	57,
	147,
	6,
	117,
	95,
	2,
	112,
	107,
	54,
	146,
	104,
	40,
	26,
	136,
	127,
	111,
	47,
	8,
	24,
	13,
	92,
	18,
	130,
	141,
	37,
	81,
	148,
	31,
	62,
	50,
	80,
	91,
	33,
	77,
	1,
	96,
	100,
	9,
	120,
	27,
	97,
	60,
	102,
	25,
	83,
	55,
	118,
	19,
	113,
	49,
	133,
	14,
	119,
	88,
	124,
	110,
	145,
	65,
	21,
	7,
	74,
	72,
	61,
	103,
	20,
	41,
	53,
	32,
	44,
	10,
	34,
	121,
	114,
	67,
	69,
	66,
	82,
	101,
	68,
	84,
	48,
	73,
	17,
	43,
	140,
}

func JoltageDiff() {
	sort.Ints(joltageAdapters)
	oneDiffsCount := 0
	// one extra hop to the device
	threeDiffsCount := 1
	if joltageAdapters[0] == 1 {
		oneDiffsCount++
	}
	if joltageAdapters[0] == 3 {
		threeDiffsCount++
	}
	for i := 1; i < len(joltageAdapters); i++ {
		diff := joltageAdapters[i] - joltageAdapters[i-1]
		if diff == 1 {
			oneDiffsCount++
		}
		if diff == 3 {
			threeDiffsCount++
		}
	}
	fmt.Printf("%#v\n", joltageAdapters)
	fmt.Println("got ", oneDiffsCount, " one diffs and ", threeDiffsCount, " three diffs")
	fmt.Println("product: ", oneDiffsCount*threeDiffsCount)
}
